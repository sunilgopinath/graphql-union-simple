// Package starwars provides a example schema and resolver based on Star Wars characters.
//
// Source: https://github.com/graphql/graphql.github.io/blob/source/site/_core/swapiSchema.js
package demo

import (
	"context"

	pb "bitbucket.org/ffxblue/golang-demo/generated/grpc"
	pbs "bitbucket.org/ffxblue/golang-demo/generated/spark"
	graphql "github.com/neelance/graphql-go"
)

var Schema = `
	schema {
		query: Query
	}
	# The query type, represents all of the entry points into our object graph
	type Query {
		search(id: String!): [SearchResult]!
	}

	type GoogleAnalytics {
		# The ID of the asset
		id: ID!
		# The name of the asset
		today: Int!
		# The author of the asset
		overall: Int!
	}

	type SparkAnalytics {
		# The ID of the asset
		id: ID!
		# The name of the asset
		current: Int!
		# The author of the asset
		trending: Boolean!
	}

	union SearchResult = GoogleAnalytics | SparkAnalytics
`

type Resolver struct {
	pb.GoogleAnalyticsClient
	pbs.SparkAnalyticsClient
}

func (r *Resolver) Search(ctx context.Context, args *struct{ ID string }) ([]*searchResultResolver, error) {
	res, err := r.GoogleAnalyticsClient.AssetStatistics(ctx, &pb.AssetRequest{Name: args.ID})
	if err != nil {
		return nil, err
	}

	sres, err := r.SparkAnalyticsClient.AssetStatistics(ctx, &pbs.AssetRequest{Name: args.ID})
	if err != nil {
		return nil, err
	}

	var l []*searchResultResolver

	l = append(l, &searchResultResolver{&googleAnalyticsResolver{res}})
	l = append(l, &searchResultResolver{&sparkAnalyticsResolver{sres}})

	return l, nil
}

func (r *searchResultResolver) ToGoogleAnalytics() (*googleAnalyticsResolver, bool) {
	res, ok := r.result.(*googleAnalyticsResolver)
	return res, ok
}

func (r *searchResultResolver) ToSparkAnalytics() (*sparkAnalyticsResolver, bool) {
	res, ok := r.result.(*sparkAnalyticsResolver)
	return res, ok
}

type searchResultResolver struct {
	result interface{}
}

type googleAnalyticsResolver struct {
	ga *pb.AssetReply
}

func (r *googleAnalyticsResolver) ID() graphql.ID {
	return graphql.ID(r.ga.Id)
}

func (r *googleAnalyticsResolver) Today() int32 {
	return r.ga.Today
}

func (r *googleAnalyticsResolver) Overall() int32 {
	return r.ga.Overall
}

type sparkAnalyticsResolver struct {
	sp *pbs.AssetReply
}

func (r *sparkAnalyticsResolver) ID() graphql.ID {
	return graphql.ID(r.sp.Id)
}

func (r *sparkAnalyticsResolver) Current() int32 {
	return r.sp.Current
}

func (r *sparkAnalyticsResolver) Trending() bool {
	return r.sp.Trending
}
