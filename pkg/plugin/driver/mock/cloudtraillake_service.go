package ctlservicemock

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

const SinglePageResponseQueryId = "singlePageResponse"
const MultiPageResponseQueryId = "multiPageResponse"

//var columnMetaData = []*cloudtrail.ColumnMetadata{
//	{
//		Name:     aws.String("col1"),
//		Nullable: aws.Int64(1),
//		TypeName: aws.String("varchar"),
//	},
//	{
//		Name:     aws.String("col2"),
//		Nullable: aws.Int64(1),
//		TypeName: aws.String("varchar"),
//	},
//}

var twoRecords = [][]map[string]*string{}

type CtlService struct {
	CalledTimesCounter   int
	CalledTimesCountDown int
}

func (s *CtlService) AddTags(input *cloudtrail.AddTagsInput) (*cloudtrail.AddTagsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) AddTagsWithContext(context aws.Context, input *cloudtrail.AddTagsInput, option ...request.Option) (*cloudtrail.AddTagsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) AddTagsRequest(input *cloudtrail.AddTagsInput) (*request.Request, *cloudtrail.AddTagsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) CreateChannel(input *cloudtrail.CreateChannelInput) (*cloudtrail.CreateChannelOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) CreateChannelWithContext(context aws.Context, input *cloudtrail.CreateChannelInput, option ...request.Option) (*cloudtrail.CreateChannelOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) CreateChannelRequest(input *cloudtrail.CreateChannelInput) (*request.Request, *cloudtrail.CreateChannelOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) CreateEventDataStore(input *cloudtrail.CreateEventDataStoreInput) (*cloudtrail.CreateEventDataStoreOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) CreateEventDataStoreWithContext(context aws.Context, input *cloudtrail.CreateEventDataStoreInput, option ...request.Option) (*cloudtrail.CreateEventDataStoreOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) CreateEventDataStoreRequest(input *cloudtrail.CreateEventDataStoreInput) (*request.Request, *cloudtrail.CreateEventDataStoreOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) CreateTrail(input *cloudtrail.CreateTrailInput) (*cloudtrail.CreateTrailOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) CreateTrailWithContext(context aws.Context, input *cloudtrail.CreateTrailInput, option ...request.Option) (*cloudtrail.CreateTrailOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) CreateTrailRequest(input *cloudtrail.CreateTrailInput) (*request.Request, *cloudtrail.CreateTrailOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DeleteChannel(input *cloudtrail.DeleteChannelInput) (*cloudtrail.DeleteChannelOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DeleteChannelWithContext(context aws.Context, input *cloudtrail.DeleteChannelInput, option ...request.Option) (*cloudtrail.DeleteChannelOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DeleteChannelRequest(input *cloudtrail.DeleteChannelInput) (*request.Request, *cloudtrail.DeleteChannelOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DeleteEventDataStore(input *cloudtrail.DeleteEventDataStoreInput) (*cloudtrail.DeleteEventDataStoreOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DeleteEventDataStoreWithContext(context aws.Context, input *cloudtrail.DeleteEventDataStoreInput, option ...request.Option) (*cloudtrail.DeleteEventDataStoreOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DeleteEventDataStoreRequest(input *cloudtrail.DeleteEventDataStoreInput) (*request.Request, *cloudtrail.DeleteEventDataStoreOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DeleteResourcePolicy(input *cloudtrail.DeleteResourcePolicyInput) (*cloudtrail.DeleteResourcePolicyOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DeleteResourcePolicyWithContext(context aws.Context, input *cloudtrail.DeleteResourcePolicyInput, option ...request.Option) (*cloudtrail.DeleteResourcePolicyOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DeleteResourcePolicyRequest(input *cloudtrail.DeleteResourcePolicyInput) (*request.Request, *cloudtrail.DeleteResourcePolicyOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DeleteTrail(input *cloudtrail.DeleteTrailInput) (*cloudtrail.DeleteTrailOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DeleteTrailWithContext(context aws.Context, input *cloudtrail.DeleteTrailInput, option ...request.Option) (*cloudtrail.DeleteTrailOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DeleteTrailRequest(input *cloudtrail.DeleteTrailInput) (*request.Request, *cloudtrail.DeleteTrailOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DeregisterOrganizationDelegatedAdmin(input *cloudtrail.DeregisterOrganizationDelegatedAdminInput) (*cloudtrail.DeregisterOrganizationDelegatedAdminOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DeregisterOrganizationDelegatedAdminWithContext(context aws.Context, input *cloudtrail.DeregisterOrganizationDelegatedAdminInput, option ...request.Option) (*cloudtrail.DeregisterOrganizationDelegatedAdminOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DeregisterOrganizationDelegatedAdminRequest(input *cloudtrail.DeregisterOrganizationDelegatedAdminInput) (*request.Request, *cloudtrail.DeregisterOrganizationDelegatedAdminOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DescribeTrails(input *cloudtrail.DescribeTrailsInput) (*cloudtrail.DescribeTrailsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DescribeTrailsWithContext(context aws.Context, input *cloudtrail.DescribeTrailsInput, option ...request.Option) (*cloudtrail.DescribeTrailsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) DescribeTrailsRequest(input *cloudtrail.DescribeTrailsInput) (*request.Request, *cloudtrail.DescribeTrailsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetChannel(input *cloudtrail.GetChannelInput) (*cloudtrail.GetChannelOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetChannelWithContext(context aws.Context, input *cloudtrail.GetChannelInput, option ...request.Option) (*cloudtrail.GetChannelOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetChannelRequest(input *cloudtrail.GetChannelInput) (*request.Request, *cloudtrail.GetChannelOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetEventDataStore(input *cloudtrail.GetEventDataStoreInput) (*cloudtrail.GetEventDataStoreOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetEventDataStoreWithContext(context aws.Context, input *cloudtrail.GetEventDataStoreInput, option ...request.Option) (*cloudtrail.GetEventDataStoreOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetEventDataStoreRequest(input *cloudtrail.GetEventDataStoreInput) (*request.Request, *cloudtrail.GetEventDataStoreOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetEventSelectors(input *cloudtrail.GetEventSelectorsInput) (*cloudtrail.GetEventSelectorsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetEventSelectorsWithContext(context aws.Context, input *cloudtrail.GetEventSelectorsInput, option ...request.Option) (*cloudtrail.GetEventSelectorsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetEventSelectorsRequest(input *cloudtrail.GetEventSelectorsInput) (*request.Request, *cloudtrail.GetEventSelectorsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetImport(input *cloudtrail.GetImportInput) (*cloudtrail.GetImportOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetImportWithContext(context aws.Context, input *cloudtrail.GetImportInput, option ...request.Option) (*cloudtrail.GetImportOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetImportRequest(input *cloudtrail.GetImportInput) (*request.Request, *cloudtrail.GetImportOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetInsightSelectors(input *cloudtrail.GetInsightSelectorsInput) (*cloudtrail.GetInsightSelectorsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetInsightSelectorsWithContext(context aws.Context, input *cloudtrail.GetInsightSelectorsInput, option ...request.Option) (*cloudtrail.GetInsightSelectorsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetInsightSelectorsRequest(input *cloudtrail.GetInsightSelectorsInput) (*request.Request, *cloudtrail.GetInsightSelectorsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetQueryResults(input *cloudtrail.GetQueryResultsInput) (*cloudtrail.GetQueryResultsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetQueryResultsWithContext(context aws.Context, input *cloudtrail.GetQueryResultsInput, option ...request.Option) (*cloudtrail.GetQueryResultsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetQueryResultsRequest(input *cloudtrail.GetQueryResultsInput) (*request.Request, *cloudtrail.GetQueryResultsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetQueryResultsPages(input *cloudtrail.GetQueryResultsInput, f func(*cloudtrail.GetQueryResultsOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetQueryResultsPagesWithContext(context aws.Context, input *cloudtrail.GetQueryResultsInput, f func(*cloudtrail.GetQueryResultsOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetResourcePolicy(input *cloudtrail.GetResourcePolicyInput) (*cloudtrail.GetResourcePolicyOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetResourcePolicyWithContext(context aws.Context, input *cloudtrail.GetResourcePolicyInput, option ...request.Option) (*cloudtrail.GetResourcePolicyOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetResourcePolicyRequest(input *cloudtrail.GetResourcePolicyInput) (*request.Request, *cloudtrail.GetResourcePolicyOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetTrail(input *cloudtrail.GetTrailInput) (*cloudtrail.GetTrailOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetTrailWithContext(context aws.Context, input *cloudtrail.GetTrailInput, option ...request.Option) (*cloudtrail.GetTrailOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetTrailRequest(input *cloudtrail.GetTrailInput) (*request.Request, *cloudtrail.GetTrailOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetTrailStatus(input *cloudtrail.GetTrailStatusInput) (*cloudtrail.GetTrailStatusOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetTrailStatusWithContext(context aws.Context, input *cloudtrail.GetTrailStatusInput, option ...request.Option) (*cloudtrail.GetTrailStatusOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) GetTrailStatusRequest(input *cloudtrail.GetTrailStatusInput) (*request.Request, *cloudtrail.GetTrailStatusOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListChannels(input *cloudtrail.ListChannelsInput) (*cloudtrail.ListChannelsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListChannelsWithContext(context aws.Context, input *cloudtrail.ListChannelsInput, option ...request.Option) (*cloudtrail.ListChannelsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListChannelsRequest(input *cloudtrail.ListChannelsInput) (*request.Request, *cloudtrail.ListChannelsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListChannelsPages(input *cloudtrail.ListChannelsInput, f func(*cloudtrail.ListChannelsOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListChannelsPagesWithContext(context aws.Context, input *cloudtrail.ListChannelsInput, f func(*cloudtrail.ListChannelsOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListEventDataStores(input *cloudtrail.ListEventDataStoresInput) (*cloudtrail.ListEventDataStoresOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListEventDataStoresWithContext(context aws.Context, input *cloudtrail.ListEventDataStoresInput, option ...request.Option) (*cloudtrail.ListEventDataStoresOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListEventDataStoresRequest(input *cloudtrail.ListEventDataStoresInput) (*request.Request, *cloudtrail.ListEventDataStoresOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListEventDataStoresPages(input *cloudtrail.ListEventDataStoresInput, f func(*cloudtrail.ListEventDataStoresOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListEventDataStoresPagesWithContext(context aws.Context, input *cloudtrail.ListEventDataStoresInput, f func(*cloudtrail.ListEventDataStoresOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListImportFailures(input *cloudtrail.ListImportFailuresInput) (*cloudtrail.ListImportFailuresOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListImportFailuresWithContext(context aws.Context, input *cloudtrail.ListImportFailuresInput, option ...request.Option) (*cloudtrail.ListImportFailuresOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListImportFailuresRequest(input *cloudtrail.ListImportFailuresInput) (*request.Request, *cloudtrail.ListImportFailuresOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListImportFailuresPages(input *cloudtrail.ListImportFailuresInput, f func(*cloudtrail.ListImportFailuresOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListImportFailuresPagesWithContext(context aws.Context, input *cloudtrail.ListImportFailuresInput, f func(*cloudtrail.ListImportFailuresOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListImports(input *cloudtrail.ListImportsInput) (*cloudtrail.ListImportsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListImportsWithContext(context aws.Context, input *cloudtrail.ListImportsInput, option ...request.Option) (*cloudtrail.ListImportsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListImportsRequest(input *cloudtrail.ListImportsInput) (*request.Request, *cloudtrail.ListImportsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListImportsPages(input *cloudtrail.ListImportsInput, f func(*cloudtrail.ListImportsOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListImportsPagesWithContext(context aws.Context, input *cloudtrail.ListImportsInput, f func(*cloudtrail.ListImportsOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListPublicKeys(input *cloudtrail.ListPublicKeysInput) (*cloudtrail.ListPublicKeysOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListPublicKeysWithContext(context aws.Context, input *cloudtrail.ListPublicKeysInput, option ...request.Option) (*cloudtrail.ListPublicKeysOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListPublicKeysRequest(input *cloudtrail.ListPublicKeysInput) (*request.Request, *cloudtrail.ListPublicKeysOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListPublicKeysPages(input *cloudtrail.ListPublicKeysInput, f func(*cloudtrail.ListPublicKeysOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListPublicKeysPagesWithContext(context aws.Context, input *cloudtrail.ListPublicKeysInput, f func(*cloudtrail.ListPublicKeysOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListQueries(input *cloudtrail.ListQueriesInput) (*cloudtrail.ListQueriesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListQueriesWithContext(context aws.Context, input *cloudtrail.ListQueriesInput, option ...request.Option) (*cloudtrail.ListQueriesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListQueriesRequest(input *cloudtrail.ListQueriesInput) (*request.Request, *cloudtrail.ListQueriesOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListQueriesPages(input *cloudtrail.ListQueriesInput, f func(*cloudtrail.ListQueriesOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListQueriesPagesWithContext(context aws.Context, input *cloudtrail.ListQueriesInput, f func(*cloudtrail.ListQueriesOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListTags(input *cloudtrail.ListTagsInput) (*cloudtrail.ListTagsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListTagsWithContext(context aws.Context, input *cloudtrail.ListTagsInput, option ...request.Option) (*cloudtrail.ListTagsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListTagsRequest(input *cloudtrail.ListTagsInput) (*request.Request, *cloudtrail.ListTagsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListTagsPages(input *cloudtrail.ListTagsInput, f func(*cloudtrail.ListTagsOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListTagsPagesWithContext(context aws.Context, input *cloudtrail.ListTagsInput, f func(*cloudtrail.ListTagsOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListTrails(input *cloudtrail.ListTrailsInput) (*cloudtrail.ListTrailsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListTrailsWithContext(context aws.Context, input *cloudtrail.ListTrailsInput, option ...request.Option) (*cloudtrail.ListTrailsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListTrailsRequest(input *cloudtrail.ListTrailsInput) (*request.Request, *cloudtrail.ListTrailsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListTrailsPages(input *cloudtrail.ListTrailsInput, f func(*cloudtrail.ListTrailsOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) ListTrailsPagesWithContext(context aws.Context, input *cloudtrail.ListTrailsInput, f func(*cloudtrail.ListTrailsOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) LookupEvents(input *cloudtrail.LookupEventsInput) (*cloudtrail.LookupEventsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) LookupEventsWithContext(context aws.Context, input *cloudtrail.LookupEventsInput, option ...request.Option) (*cloudtrail.LookupEventsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) LookupEventsRequest(input *cloudtrail.LookupEventsInput) (*request.Request, *cloudtrail.LookupEventsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) LookupEventsPages(input *cloudtrail.LookupEventsInput, f func(*cloudtrail.LookupEventsOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) LookupEventsPagesWithContext(context aws.Context, input *cloudtrail.LookupEventsInput, f func(*cloudtrail.LookupEventsOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) PutEventSelectors(input *cloudtrail.PutEventSelectorsInput) (*cloudtrail.PutEventSelectorsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) PutEventSelectorsWithContext(context aws.Context, input *cloudtrail.PutEventSelectorsInput, option ...request.Option) (*cloudtrail.PutEventSelectorsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) PutEventSelectorsRequest(input *cloudtrail.PutEventSelectorsInput) (*request.Request, *cloudtrail.PutEventSelectorsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) PutInsightSelectors(input *cloudtrail.PutInsightSelectorsInput) (*cloudtrail.PutInsightSelectorsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) PutInsightSelectorsWithContext(context aws.Context, input *cloudtrail.PutInsightSelectorsInput, option ...request.Option) (*cloudtrail.PutInsightSelectorsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) PutInsightSelectorsRequest(input *cloudtrail.PutInsightSelectorsInput) (*request.Request, *cloudtrail.PutInsightSelectorsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) PutResourcePolicy(input *cloudtrail.PutResourcePolicyInput) (*cloudtrail.PutResourcePolicyOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) PutResourcePolicyWithContext(context aws.Context, input *cloudtrail.PutResourcePolicyInput, option ...request.Option) (*cloudtrail.PutResourcePolicyOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) PutResourcePolicyRequest(input *cloudtrail.PutResourcePolicyInput) (*request.Request, *cloudtrail.PutResourcePolicyOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) RegisterOrganizationDelegatedAdmin(input *cloudtrail.RegisterOrganizationDelegatedAdminInput) (*cloudtrail.RegisterOrganizationDelegatedAdminOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) RegisterOrganizationDelegatedAdminWithContext(context aws.Context, input *cloudtrail.RegisterOrganizationDelegatedAdminInput, option ...request.Option) (*cloudtrail.RegisterOrganizationDelegatedAdminOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) RegisterOrganizationDelegatedAdminRequest(input *cloudtrail.RegisterOrganizationDelegatedAdminInput) (*request.Request, *cloudtrail.RegisterOrganizationDelegatedAdminOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) RemoveTags(input *cloudtrail.RemoveTagsInput) (*cloudtrail.RemoveTagsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) RemoveTagsWithContext(context aws.Context, input *cloudtrail.RemoveTagsInput, option ...request.Option) (*cloudtrail.RemoveTagsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) RemoveTagsRequest(input *cloudtrail.RemoveTagsInput) (*request.Request, *cloudtrail.RemoveTagsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) RestoreEventDataStore(input *cloudtrail.RestoreEventDataStoreInput) (*cloudtrail.RestoreEventDataStoreOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) RestoreEventDataStoreWithContext(context aws.Context, input *cloudtrail.RestoreEventDataStoreInput, option ...request.Option) (*cloudtrail.RestoreEventDataStoreOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) RestoreEventDataStoreRequest(input *cloudtrail.RestoreEventDataStoreInput) (*request.Request, *cloudtrail.RestoreEventDataStoreOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StartEventDataStoreIngestion(input *cloudtrail.StartEventDataStoreIngestionInput) (*cloudtrail.StartEventDataStoreIngestionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StartEventDataStoreIngestionWithContext(context aws.Context, input *cloudtrail.StartEventDataStoreIngestionInput, option ...request.Option) (*cloudtrail.StartEventDataStoreIngestionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StartEventDataStoreIngestionRequest(input *cloudtrail.StartEventDataStoreIngestionInput) (*request.Request, *cloudtrail.StartEventDataStoreIngestionOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StartImport(input *cloudtrail.StartImportInput) (*cloudtrail.StartImportOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StartImportWithContext(context aws.Context, input *cloudtrail.StartImportInput, option ...request.Option) (*cloudtrail.StartImportOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StartImportRequest(input *cloudtrail.StartImportInput) (*request.Request, *cloudtrail.StartImportOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StartLogging(input *cloudtrail.StartLoggingInput) (*cloudtrail.StartLoggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StartLoggingWithContext(context aws.Context, input *cloudtrail.StartLoggingInput, option ...request.Option) (*cloudtrail.StartLoggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StartLoggingRequest(input *cloudtrail.StartLoggingInput) (*request.Request, *cloudtrail.StartLoggingOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StartQuery(input *cloudtrail.StartQueryInput) (*cloudtrail.StartQueryOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StartQueryWithContext(context aws.Context, input *cloudtrail.StartQueryInput, option ...request.Option) (*cloudtrail.StartQueryOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StartQueryRequest(input *cloudtrail.StartQueryInput) (*request.Request, *cloudtrail.StartQueryOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StopEventDataStoreIngestion(input *cloudtrail.StopEventDataStoreIngestionInput) (*cloudtrail.StopEventDataStoreIngestionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StopEventDataStoreIngestionWithContext(context aws.Context, input *cloudtrail.StopEventDataStoreIngestionInput, option ...request.Option) (*cloudtrail.StopEventDataStoreIngestionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StopEventDataStoreIngestionRequest(input *cloudtrail.StopEventDataStoreIngestionInput) (*request.Request, *cloudtrail.StopEventDataStoreIngestionOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StopImport(input *cloudtrail.StopImportInput) (*cloudtrail.StopImportOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StopImportWithContext(context aws.Context, input *cloudtrail.StopImportInput, option ...request.Option) (*cloudtrail.StopImportOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StopImportRequest(input *cloudtrail.StopImportInput) (*request.Request, *cloudtrail.StopImportOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StopLogging(input *cloudtrail.StopLoggingInput) (*cloudtrail.StopLoggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StopLoggingWithContext(context aws.Context, input *cloudtrail.StopLoggingInput, option ...request.Option) (*cloudtrail.StopLoggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) StopLoggingRequest(input *cloudtrail.StopLoggingInput) (*request.Request, *cloudtrail.StopLoggingOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) UpdateChannel(input *cloudtrail.UpdateChannelInput) (*cloudtrail.UpdateChannelOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) UpdateChannelWithContext(context aws.Context, input *cloudtrail.UpdateChannelInput, option ...request.Option) (*cloudtrail.UpdateChannelOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) UpdateChannelRequest(input *cloudtrail.UpdateChannelInput) (*request.Request, *cloudtrail.UpdateChannelOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) UpdateEventDataStore(input *cloudtrail.UpdateEventDataStoreInput) (*cloudtrail.UpdateEventDataStoreOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) UpdateEventDataStoreWithContext(context aws.Context, input *cloudtrail.UpdateEventDataStoreInput, option ...request.Option) (*cloudtrail.UpdateEventDataStoreOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) UpdateEventDataStoreRequest(input *cloudtrail.UpdateEventDataStoreInput) (*request.Request, *cloudtrail.UpdateEventDataStoreOutput) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) UpdateTrail(input *cloudtrail.UpdateTrailInput) (*cloudtrail.UpdateTrailOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) UpdateTrailWithContext(context aws.Context, input *cloudtrail.UpdateTrailInput, option ...request.Option) (*cloudtrail.UpdateTrailOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CtlService) UpdateTrailRequest(input *cloudtrail.UpdateTrailInput) (*request.Request, *cloudtrail.UpdateTrailOutput) {
	//TODO implement me
	panic("implement me")
}

func NewMockCtlService() *CtlService {
	return &CtlService{CalledTimesCounter: 0}
}

// GetQueryResult returns a GetQueryResultsOutput
// When mockCtlService.calledTimesCountDown is more than 0, the GetQueryResultsOutput will have a next token
func (s *CtlService) GetQueryResult(input *cloudtrail.GetQueryResultsInput) (*cloudtrail.GetQueryResultsOutput, error) {
	s.CalledTimesCountDown--
	s.CalledTimesCounter++

	if s.CalledTimesCountDown == 0 {
		return &cloudtrail.GetQueryResultsOutput{
			// FIXME fill
			QueryResultRows: [][]map[string]*string{},
		}, nil
	}

	return &cloudtrail.GetQueryResultsOutput{
		// Records:        twoRecords,
		// FIXME fill
		QueryResultRows: [][]map[string]*string{},
		NextToken:       aws.String("nexttoken"),
	}, nil
}

const DESCRIBE_STATEMENT_FAILED = "DESCRIBE_STATEMENT_FAILED"
const DESCRIBE_STATEMENT_SUCCEEDED = "DESCRIBE_STATEMENT_FINISHED"

// DescribeQueryWithContext returns a DescribeQueryOutput
// When DescribeQueryInput.Id == DESCRIBE_STATEMENT_FAILED, an the output will include an error message that is equal to the input id
// When DescribeQueryInput.Id == DESCRIBE_STATEMENT_FINISHED, the output will have a status cloudtrail.StatusStringFinished once mockCtlService.calledTimesCountDown == 0
func (s *CtlService) DescribeQueryWithContext(_ aws.Context, input *cloudtrail.DescribeQueryInput, _ ...request.Option) (*cloudtrail.DescribeQueryOutput, error) {
	s.CalledTimesCountDown--
	s.CalledTimesCounter++

	output := &cloudtrail.DescribeQueryOutput{}
	if s.CalledTimesCountDown == 0 {
		if *input.QueryId == DESCRIBE_STATEMENT_FAILED {
			output.QueryStatus = aws.String(cloudtrail.QueryStatusFailed)
			output.ErrorMessage = aws.String(DESCRIBE_STATEMENT_FAILED)
		} else {
			output.QueryStatus = aws.String(cloudtrail.QueryStatusFinished)
		}
	} else {
		output.QueryStatus = aws.String(cloudtrail.QueryStatusRunning)
	}
	return output, nil
}

func (s *CtlService) CancelQuery(*cloudtrail.CancelQueryInput) (*cloudtrail.CancelQueryOutput, error) {
	panic("not implemented")
}

func (s *CtlService) CancelQueryWithContext(aws.Context, *cloudtrail.CancelQueryInput, ...request.Option) (*cloudtrail.CancelQueryOutput, error) {
	panic("not implemented")
}

func (s *CtlService) CancelQueryRequest(*cloudtrail.CancelQueryInput) (*request.Request, *cloudtrail.CancelQueryOutput) {
	panic("not implemented")
}

func (s *CtlService) DescribeQuery(*cloudtrail.DescribeQueryInput) (*cloudtrail.DescribeQueryOutput, error) {
	panic("not implemented")
}

func (s *CtlService) DescribeQueryRequest(*cloudtrail.DescribeQueryInput) (*request.Request, *cloudtrail.DescribeQueryOutput) {
	panic("not implemented")
}

//func (s *CtlService) DescribeTable(*cloudtrail.DescribeTableInput) (*cloudtrail.DescribeTableOutput, error) {
//	panic("not implemented")
//}

//func (s *CtlService) DescribeTableWithContext(aws.Context, *cloudtrail.DescribeTableInput, ...request.Option) (*cloudtrail.DescribeTableOutput, error) {
//	panic("not implemented")
//}
//
//func (s *CtlService) DescribeTableRequest(*cloudtrail.DescribeTableInput) (*request.Request, *cloudtrail.DescribeTableOutput) {
//	panic("not implemented")
//}
//
//func (s *CtlService) DescribeTablePages(*cloudtrail.DescribeTableInput, func(*cloudtrail.DescribeTableOutput, bool) bool) error {
//	panic("not implemented")
//}
//
//func (s *CtlService) DescribeTablePagesWithContext(aws.Context, *cloudtrail.DescribeTableInput, func(*cloudtrail.DescribeTableOutput, bool) bool, ...request.Option) error {
//	panic("not implemented")
//}

func (s *CtlService) ExecuteQuery(input *cloudtrail.StartQueryInput) (*cloudtrail.StartQueryOutput, error) {
	panic("not implemented")
}

func (s *CtlService) ExecuteQueryWithContext(aws.Context, *cloudtrail.StartQueryInput, ...request.Option) (*cloudtrail.StartQueryOutput, error) {
	panic("not implemented")
}

func (s *CtlService) ExecuteQueryRequest(*cloudtrail.StartQueryInput) (*request.Request, *cloudtrail.StartQueryOutput) {
	panic("not implemented")
}

func (s *CtlService) GetQueryResultWithContext(aws.Context, *cloudtrail.GetQueryResultsInput, ...request.Option) (*cloudtrail.GetQueryResultsOutput, error) {
	panic("not implemented")
}

func (s *CtlService) GetQueryResultRequest(*cloudtrail.GetQueryResultsInput) (*request.Request, *cloudtrail.GetQueryResultsOutput) {
	panic("not implemented")
}

func (s *CtlService) GetQueryResultPages(*cloudtrail.GetQueryResultsInput, func(*cloudtrail.GetQueryResultsOutput, bool) bool) error {
	panic("not implemented")
}

func (s *CtlService) GetQueryResultPagesWithContext(aws.Context, *cloudtrail.GetQueryResultsInput, func(*cloudtrail.GetQueryResultsOutput, bool) bool, ...request.Option) error {
	panic("not implemented")
}

//func (s *CtlService) ListDatabases(*cloudtrail.ListDatabasesInput) (*cloudtrail.ListDatabasesOutput, error) {
//	panic("not implemented")
//}
//
//func (s *CtlService) ListDatabasesWithContext(aws.Context, *cloudtrail.ListDatabasesInput, ...request.Option) (*cloudtrail.ListDatabasesOutput, error) {
//	panic("not implemented")
//}
//
//func (s *CtlService) ListDatabasesPagesWithContext(aws.Context, *cloudtrail.ListDatabasesInput, func(*cloudtrail.ListDatabasesOutput, bool) bool, ...request.Option) error {
//	panic("not implemented")
//}
//
//func (s *CtlService) ListDatabasesRequest(*cloudtrail.ListDatabasesInput) (*request.Request, *cloudtrail.ListDatabasesOutput) {
//	panic("not implemented")
//}
//
//func (s *CtlService) ListDatabasesPages(*cloudtrail.ListDatabasesInput, func(*cloudtrail.ListDatabasesOutput, bool) bool) error {
//	panic("not implemented")
//}
//
//func (s *CtlService) ListSchemas(*cloudtrail.ListSchemasInput) (*cloudtrail.ListSchemasOutput, error) {
//	panic("not implemented")
//}
//
//func (s *CtlService) ListSchemasWithContext(aws.Context, *cloudtrail.ListSchemasInput, ...request.Option) (*cloudtrail.ListSchemasOutput, error) {
//	panic("not implemented")
//}
//
//func (s *CtlService) ListSchemasRequest(*cloudtrail.ListSchemasInput) (*request.Request, *cloudtrail.ListSchemasOutput) {
//	panic("not implemented")
//}
//
//func (s *CtlService) ListSchemasPages(*cloudtrail.ListSchemasInput, func(*cloudtrail.ListSchemasOutput, bool) bool) error {
//	panic("not implemented")
//}
//
//func (s *CtlService) ListSchemasPagesWithContext(aws.Context, *cloudtrail.ListSchemasInput, func(*cloudtrail.ListSchemasOutput, bool) bool, ...request.Option) error {
//	panic("not implemented")
//}

func (s *CtlService) ListQuerys(input *cloudtrail.ListQueriesInput) (*cloudtrail.ListQueriesOutput, error) {
	panic("not implemented")
}

func (s *CtlService) ListQuerysWithContext(aws.Context, *cloudtrail.ListQueriesInput, ...request.Option) (*cloudtrail.ListQueriesOutput, error) {
	panic("not implemented")
}

func (s *CtlService) ListQuerysRequest(*cloudtrail.ListQueriesInput) (*request.Request, *cloudtrail.ListQueriesOutput) {
	panic("not implemented")
}

func (s *CtlService) ListQuerysPages(*cloudtrail.ListQueriesInput, func(*cloudtrail.ListQueriesOutput, bool) bool) error {
	panic("not implemented")
}

func (s *CtlService) ListQuerysPagesWithContext(aws.Context, *cloudtrail.ListQueriesInput, func(*cloudtrail.ListQueriesOutput, bool) bool, ...request.Option) error {
	panic("not implemented")
}

//func (s *CtlService) ListTables(*cloudtrail.ListTablesInput) (*cloudtrail.ListTablesOutput, error) {
//	panic("not implemented")
//}

//func (s *CtlService) ListTablesWithContext(aws.Context, *cloudtrail.ListTablesInput, ...request.Option) (*cloudtrail.ListTablesOutput, error) {
//	panic("not implemented")
//}
//
//func (s *CtlService) ListTablesRequest(*cloudtrail.ListTablesInput) (*request.Request, *cloudtrail.ListTablesOutput) {
//	panic("not implemented")
//}
//
//func (s *CtlService) ListTablesPages(*cloudtrail.ListTablesInput, func(*cloudtrail.ListTablesOutput, bool) bool) error {
//	panic("not implemented")
//}
//
//func (s *CtlService) ListTablesPagesWithContext(aws.Context, *cloudtrail.ListTablesInput, func(*cloudtrail.ListTablesOutput, bool) bool, ...request.Option) error {
//	panic("not implemented")
//}
//
//func (s *CtlService) BatchExecuteQuery(input *cloudtrail.BatchExecuteQueryInput) (*cloudtrail.BatchExecuteQueryOutput, error) {
//	panic("not implemented")
//}
//
//func (s *CtlService) BatchExecuteQueryWithContext(context aws.Context, input *cloudtrail.BatchExecuteQueryInput, option ...request.Option) (*cloudtrail.BatchExecuteQueryOutput, error) {
//	panic("not implemented")
//}
//
//func (s *CtlService) BatchExecuteQueryRequest(input *cloudtrail.BatchExecuteQueryInput) (*request.Request, *cloudtrail.BatchExecuteQueryOutput) {
//	panic("not implemented")
//}
