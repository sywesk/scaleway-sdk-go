// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package container provides methods and message types of the container v1beta1 API.
package container

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ fmt.Stringer
	_ json.Unmarshaler
	_ url.URL
	_ net.IP
	_ http.Header
	_ bytes.Reader
	_ time.Time
	_ = strings.Join

	_ scw.ScalewayRequest
	_ marshaler.Duration
	_ scw.File
	_ = parameter.AddToQuery
	_ = namegenerator.GetRandomName
)

// API: serverless Containers API
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type ContainerPrivacy string

const (
	// ContainerPrivacyUnknownPrivacy is [insert doc].
	ContainerPrivacyUnknownPrivacy = ContainerPrivacy("unknown_privacy")
	// ContainerPrivacyPublic is [insert doc].
	ContainerPrivacyPublic = ContainerPrivacy("public")
	// ContainerPrivacyPrivate is [insert doc].
	ContainerPrivacyPrivate = ContainerPrivacy("private")
)

func (enum ContainerPrivacy) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_privacy"
	}
	return string(enum)
}

func (enum ContainerPrivacy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContainerPrivacy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContainerPrivacy(ContainerPrivacy(tmp).String())
	return nil
}

type ContainerProtocol string

const (
	// ContainerProtocolUnknownProtocol is [insert doc].
	ContainerProtocolUnknownProtocol = ContainerProtocol("unknown_protocol")
	// ContainerProtocolHTTP1 is [insert doc].
	ContainerProtocolHTTP1 = ContainerProtocol("http1")
	// ContainerProtocolH2c is [insert doc].
	ContainerProtocolH2c = ContainerProtocol("h2c")
)

func (enum ContainerProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_protocol"
	}
	return string(enum)
}

func (enum ContainerProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContainerProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContainerProtocol(ContainerProtocol(tmp).String())
	return nil
}

type ContainerStatus string

const (
	// ContainerStatusUnknown is [insert doc].
	ContainerStatusUnknown = ContainerStatus("unknown")
	// ContainerStatusReady is [insert doc].
	ContainerStatusReady = ContainerStatus("ready")
	// ContainerStatusDeleting is [insert doc].
	ContainerStatusDeleting = ContainerStatus("deleting")
	// ContainerStatusError is [insert doc].
	ContainerStatusError = ContainerStatus("error")
	// ContainerStatusLocked is [insert doc].
	ContainerStatusLocked = ContainerStatus("locked")
	// ContainerStatusCreating is [insert doc].
	ContainerStatusCreating = ContainerStatus("creating")
	// ContainerStatusPending is [insert doc].
	ContainerStatusPending = ContainerStatus("pending")
	// ContainerStatusCreated is [insert doc].
	ContainerStatusCreated = ContainerStatus("created")
)

func (enum ContainerStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ContainerStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContainerStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContainerStatus(ContainerStatus(tmp).String())
	return nil
}

type CronStatus string

const (
	// CronStatusUnknown is [insert doc].
	CronStatusUnknown = CronStatus("unknown")
	// CronStatusReady is [insert doc].
	CronStatusReady = CronStatus("ready")
	// CronStatusDeleting is [insert doc].
	CronStatusDeleting = CronStatus("deleting")
	// CronStatusError is [insert doc].
	CronStatusError = CronStatus("error")
	// CronStatusLocked is [insert doc].
	CronStatusLocked = CronStatus("locked")
	// CronStatusCreating is [insert doc].
	CronStatusCreating = CronStatus("creating")
	// CronStatusPending is [insert doc].
	CronStatusPending = CronStatus("pending")
)

func (enum CronStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum CronStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CronStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CronStatus(CronStatus(tmp).String())
	return nil
}

type DomainStatus string

const (
	// DomainStatusUnknown is [insert doc].
	DomainStatusUnknown = DomainStatus("unknown")
	// DomainStatusReady is [insert doc].
	DomainStatusReady = DomainStatus("ready")
	// DomainStatusDeleting is [insert doc].
	DomainStatusDeleting = DomainStatus("deleting")
	// DomainStatusError is [insert doc].
	DomainStatusError = DomainStatus("error")
	// DomainStatusCreating is [insert doc].
	DomainStatusCreating = DomainStatus("creating")
	// DomainStatusPending is [insert doc].
	DomainStatusPending = DomainStatus("pending")
)

func (enum DomainStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum DomainStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainStatus(DomainStatus(tmp).String())
	return nil
}

type ListContainersRequestOrderBy string

const (
	// ListContainersRequestOrderByCreatedAtAsc is [insert doc].
	ListContainersRequestOrderByCreatedAtAsc = ListContainersRequestOrderBy("created_at_asc")
	// ListContainersRequestOrderByCreatedAtDesc is [insert doc].
	ListContainersRequestOrderByCreatedAtDesc = ListContainersRequestOrderBy("created_at_desc")
	// ListContainersRequestOrderByNameAsc is [insert doc].
	ListContainersRequestOrderByNameAsc = ListContainersRequestOrderBy("name_asc")
	// ListContainersRequestOrderByNameDesc is [insert doc].
	ListContainersRequestOrderByNameDesc = ListContainersRequestOrderBy("name_desc")
)

func (enum ListContainersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListContainersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListContainersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListContainersRequestOrderBy(ListContainersRequestOrderBy(tmp).String())
	return nil
}

type ListCronsRequestOrderBy string

const (
	// ListCronsRequestOrderByCreatedAtAsc is [insert doc].
	ListCronsRequestOrderByCreatedAtAsc = ListCronsRequestOrderBy("created_at_asc")
	// ListCronsRequestOrderByCreatedAtDesc is [insert doc].
	ListCronsRequestOrderByCreatedAtDesc = ListCronsRequestOrderBy("created_at_desc")
)

func (enum ListCronsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListCronsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCronsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCronsRequestOrderBy(ListCronsRequestOrderBy(tmp).String())
	return nil
}

type ListDomainsRequestOrderBy string

const (
	// ListDomainsRequestOrderByCreatedAtAsc is [insert doc].
	ListDomainsRequestOrderByCreatedAtAsc = ListDomainsRequestOrderBy("created_at_asc")
	// ListDomainsRequestOrderByCreatedAtDesc is [insert doc].
	ListDomainsRequestOrderByCreatedAtDesc = ListDomainsRequestOrderBy("created_at_desc")
	// ListDomainsRequestOrderByHostnameAsc is [insert doc].
	ListDomainsRequestOrderByHostnameAsc = ListDomainsRequestOrderBy("hostname_asc")
	// ListDomainsRequestOrderByHostnameDesc is [insert doc].
	ListDomainsRequestOrderByHostnameDesc = ListDomainsRequestOrderBy("hostname_desc")
)

func (enum ListDomainsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListDomainsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDomainsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDomainsRequestOrderBy(ListDomainsRequestOrderBy(tmp).String())
	return nil
}

type ListLogsRequestOrderBy string

const (
	// ListLogsRequestOrderByTimestampDesc is [insert doc].
	ListLogsRequestOrderByTimestampDesc = ListLogsRequestOrderBy("timestamp_desc")
	// ListLogsRequestOrderByTimestampAsc is [insert doc].
	ListLogsRequestOrderByTimestampAsc = ListLogsRequestOrderBy("timestamp_asc")
)

func (enum ListLogsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "timestamp_desc"
	}
	return string(enum)
}

func (enum ListLogsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListLogsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListLogsRequestOrderBy(ListLogsRequestOrderBy(tmp).String())
	return nil
}

type ListNamespacesRequestOrderBy string

const (
	// ListNamespacesRequestOrderByCreatedAtAsc is [insert doc].
	ListNamespacesRequestOrderByCreatedAtAsc = ListNamespacesRequestOrderBy("created_at_asc")
	// ListNamespacesRequestOrderByCreatedAtDesc is [insert doc].
	ListNamespacesRequestOrderByCreatedAtDesc = ListNamespacesRequestOrderBy("created_at_desc")
	// ListNamespacesRequestOrderByNameAsc is [insert doc].
	ListNamespacesRequestOrderByNameAsc = ListNamespacesRequestOrderBy("name_asc")
	// ListNamespacesRequestOrderByNameDesc is [insert doc].
	ListNamespacesRequestOrderByNameDesc = ListNamespacesRequestOrderBy("name_desc")
)

func (enum ListNamespacesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListNamespacesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNamespacesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNamespacesRequestOrderBy(ListNamespacesRequestOrderBy(tmp).String())
	return nil
}

type NamespaceStatus string

const (
	// NamespaceStatusUnknown is [insert doc].
	NamespaceStatusUnknown = NamespaceStatus("unknown")
	// NamespaceStatusReady is [insert doc].
	NamespaceStatusReady = NamespaceStatus("ready")
	// NamespaceStatusDeleting is [insert doc].
	NamespaceStatusDeleting = NamespaceStatus("deleting")
	// NamespaceStatusError is [insert doc].
	NamespaceStatusError = NamespaceStatus("error")
	// NamespaceStatusLocked is [insert doc].
	NamespaceStatusLocked = NamespaceStatus("locked")
	// NamespaceStatusCreating is [insert doc].
	NamespaceStatusCreating = NamespaceStatus("creating")
	// NamespaceStatusPending is [insert doc].
	NamespaceStatusPending = NamespaceStatus("pending")
)

func (enum NamespaceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum NamespaceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NamespaceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NamespaceStatus(NamespaceStatus(tmp).String())
	return nil
}

type NullValue string

const (
	// NullValueNULLVALUE is [insert doc].
	NullValueNULLVALUE = NullValue("NULL_VALUE")
)

func (enum NullValue) String() string {
	if enum == "" {
		// return default value if empty
		return "NULL_VALUE"
	}
	return string(enum)
}

func (enum NullValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NullValue) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NullValue(NullValue(tmp).String())
	return nil
}

// Container: container
type Container struct {
	ID string `json:"id"`

	Name string `json:"name"`

	NamespaceID string `json:"namespace_id"`
	// Status:
	//
	// Default value: unknown
	Status ContainerStatus `json:"status"`

	EnvironmentVariables map[string]string `json:"environment_variables"`

	MinScale uint32 `json:"min_scale"`

	MaxScale uint32 `json:"max_scale"`

	MemoryLimit uint32 `json:"memory_limit"`

	CPULimit uint32 `json:"cpu_limit"`

	Timeout *scw.Duration `json:"timeout"`

	ErrorMessage *string `json:"error_message"`
	// Privacy:
	//
	// Default value: unknown_privacy
	Privacy ContainerPrivacy `json:"privacy"`

	Description *string `json:"description"`

	RegistryImage string `json:"registry_image"`

	MaxConcurrency uint32 `json:"max_concurrency"`

	DomainName string `json:"domain_name"`
	// Protocol:
	//
	// Default value: unknown_protocol
	Protocol ContainerProtocol `json:"protocol"`

	Port uint32 `json:"port"`

	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`

	Region scw.Region `json:"region"`
}

// Cron: cron
type Cron struct {
	ID string `json:"id"`

	ContainerID string `json:"container_id"`

	Schedule string `json:"schedule"`

	Args []byte `json:"args"`
	// Status:
	//
	// Default value: unknown
	Status CronStatus `json:"status"`
}

// Domain: domain
type Domain struct {
	ID string `json:"id"`

	Hostname string `json:"hostname"`

	ContainerID string `json:"container_id"`

	URL string `json:"url"`
	// Status:
	//
	// Default value: unknown
	Status DomainStatus `json:"status"`

	ErrorMessage *string `json:"error_message"`
}

// ListContainersResponse: list containers response
type ListContainersResponse struct {
	Containers []*Container `json:"containers"`

	TotalCount uint32 `json:"total_count"`
}

// ListCronsResponse: list crons response
type ListCronsResponse struct {
	Crons []*Cron `json:"crons"`

	TotalCount uint32 `json:"total_count"`
}

// ListDomainsResponse: list domains response
type ListDomainsResponse struct {
	Domains []*Domain `json:"domains"`

	TotalCount uint32 `json:"total_count"`
}

// ListLogsResponse: list logs response
type ListLogsResponse struct {
	Logs []*Log `json:"logs"`

	TotalCount uint32 `json:"total_count"`
}

// ListNamespacesResponse: list namespaces response
type ListNamespacesResponse struct {
	Namespaces []*Namespace `json:"namespaces"`

	TotalCount uint32 `json:"total_count"`
}

// Log: log
type Log struct {
	Message string `json:"message"`

	Timestamp *time.Time `json:"timestamp"`

	ID string `json:"id"`
}

// Namespace: namespace
type Namespace struct {
	ID string `json:"id"`

	Name string `json:"name"`

	EnvironmentVariables map[string]string `json:"environment_variables"`

	OrganizationID string `json:"organization_id"`

	ProjectID string `json:"project_id"`
	// Status:
	//
	// Default value: unknown
	Status NamespaceStatus `json:"status"`

	RegistryNamespaceID string `json:"registry_namespace_id"`

	ErrorMessage *string `json:"error_message"`

	RegistryEndpoint string `json:"registry_endpoint"`

	Description *string `json:"description"`

	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`

	Region scw.Region `json:"region"`
}

type Secret struct {
	Key string `json:"key"`

	Value *string `json:"value"`
}

type Token struct {
	Token string `json:"token"`

	PublicKey string `json:"public_key"`
}

// Service API

type ListNamespacesRequest struct {
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListNamespacesRequestOrderBy `json:"-"`

	Name *string `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`
}

// ListNamespaces: list all your namespaces
func (s *API) ListNamespaces(req *ListNamespacesRequest, opts ...scw.RequestOption) (*ListNamespacesResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListNamespacesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetNamespaceRequest struct {
	Region scw.Region `json:"-"`

	NamespaceID string `json:"-"`
}

// GetNamespace: get a namespace
//
// Get the namespace associated with the given id.
func (s *API) GetNamespace(req *GetNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NamespaceID) == "" {
		return nil, errors.New("field NamespaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
		Headers: http.Header{},
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateNamespaceRequest struct {
	Region scw.Region `json:"-"`

	Name string `json:"name"`

	EnvironmentVariables *map[string]string `json:"environment_variables"`

	ProjectID string `json:"project_id"`

	Description *string `json:"description"`

	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`
}

// CreateNamespace: create a new namespace
func (s *API) CreateNamespace(req *CreateNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("cns")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateNamespaceRequest struct {
	Region scw.Region `json:"-"`

	NamespaceID string `json:"-"`

	EnvironmentVariables *map[string]string `json:"environment_variables"`

	Description *string `json:"description"`

	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`
}

// UpdateNamespace: update an existing namespace
//
// Update the space associated with the given id.
func (s *API) UpdateNamespace(req *UpdateNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NamespaceID) == "" {
		return nil, errors.New("field NamespaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteNamespaceRequest struct {
	Region scw.Region `json:"-"`

	NamespaceID string `json:"-"`
}

// DeleteNamespace: delete an existing namespace
//
// Delete the namespace associated with the given id.
func (s *API) DeleteNamespace(req *DeleteNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NamespaceID) == "" {
		return nil, errors.New("field NamespaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
		Headers: http.Header{},
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListContainersRequest struct {
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListContainersRequestOrderBy `json:"-"`

	NamespaceID string `json:"-"`

	Name *string `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`
}

// ListContainers: list all your containers
func (s *API) ListContainers(req *ListContainersRequest, opts ...scw.RequestOption) (*ListContainersResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListContainersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetContainerRequest struct {
	Region scw.Region `json:"-"`

	ContainerID string `json:"-"`
}

// GetContainer: get a container
//
// Get the container associated with the given id.
func (s *API) GetContainer(req *GetContainerRequest, opts ...scw.RequestOption) (*Container, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ContainerID) == "" {
		return nil, errors.New("field ContainerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "",
		Headers: http.Header{},
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateContainerRequest struct {
	Region scw.Region `json:"-"`

	NamespaceID string `json:"namespace_id"`

	Name string `json:"name"`

	EnvironmentVariables *map[string]string `json:"environment_variables"`

	MinScale *uint32 `json:"min_scale"`

	MaxScale *uint32 `json:"max_scale"`

	MemoryLimit *uint32 `json:"memory_limit"`

	Timeout *scw.Duration `json:"timeout"`
	// Privacy:
	//
	// Default value: unknown_privacy
	Privacy ContainerPrivacy `json:"privacy"`

	Description *string `json:"description"`

	RegistryImage *string `json:"registry_image"`

	MaxConcurrency *uint32 `json:"max_concurrency"`

	DomainName *string `json:"domain_name"`
	// Protocol:
	//
	// Default value: unknown_protocol
	Protocol ContainerProtocol `json:"protocol"`

	Port *uint32 `json:"port"`

	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`
}

// CreateContainer: create a new container
func (s *API) CreateContainer(req *CreateContainerRequest, opts ...scw.RequestOption) (*Container, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("ctnr")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateContainerRequest struct {
	Region scw.Region `json:"-"`

	ContainerID string `json:"-"`

	EnvironmentVariables *map[string]string `json:"environment_variables"`

	MinScale *uint32 `json:"min_scale"`

	MaxScale *uint32 `json:"max_scale"`

	MemoryLimit *uint32 `json:"memory_limit"`

	Timeout *scw.Duration `json:"timeout"`

	Redeploy *bool `json:"redeploy"`
	// Privacy:
	//
	// Default value: unknown_privacy
	Privacy ContainerPrivacy `json:"privacy"`

	Description *string `json:"description"`

	RegistryImage *string `json:"registry_image"`

	MaxConcurrency *uint32 `json:"max_concurrency"`

	DomainName *string `json:"domain_name"`
	// Protocol:
	//
	// Default value: unknown_protocol
	Protocol ContainerProtocol `json:"protocol"`

	Port *uint32 `json:"port"`

	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`
}

// UpdateContainer: update an existing container
//
// Update the container associated with the given id.
func (s *API) UpdateContainer(req *UpdateContainerRequest, opts ...scw.RequestOption) (*Container, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ContainerID) == "" {
		return nil, errors.New("field ContainerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteContainerRequest struct {
	Region scw.Region `json:"-"`

	ContainerID string `json:"-"`
}

// DeleteContainer: delete a container
//
// Delete the container associated with the given id.
func (s *API) DeleteContainer(req *DeleteContainerRequest, opts ...scw.RequestOption) (*Container, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ContainerID) == "" {
		return nil, errors.New("field ContainerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "",
		Headers: http.Header{},
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeployContainerRequest struct {
	Region scw.Region `json:"-"`

	ContainerID string `json:"-"`
}

// DeployContainer: deploy a container
//
// Deploy a container associated with the given id.
func (s *API) DeployContainer(req *DeployContainerRequest, opts ...scw.RequestOption) (*Container, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ContainerID) == "" {
		return nil, errors.New("field ContainerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "/deploy",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListCronsRequest struct {
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListCronsRequestOrderBy `json:"-"`

	ContainerID string `json:"-"`
}

// ListCrons: list all your crons
func (s *API) ListCrons(req *ListCronsRequest, opts ...scw.RequestOption) (*ListCronsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "container_id", req.ContainerID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListCronsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetCronRequest struct {
	Region scw.Region `json:"-"`

	CronID string `json:"-"`
}

// GetCron: get a cron
//
// Get the cron associated with the given id.
func (s *API) GetCron(req *GetCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CronID) == "" {
		return nil, errors.New("field CronID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
		Headers: http.Header{},
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateCronRequest struct {
	Region scw.Region `json:"-"`

	ContainerID string `json:"container_id"`

	Schedule string `json:"schedule"`

	Args []byte `json:"args"`
}

// CreateCron: create a new cron
func (s *API) CreateCron(req *CreateCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateCronRequest struct {
	Region scw.Region `json:"-"`

	CronID string `json:"-"`

	ContainerID *string `json:"container_id"`

	Schedule *string `json:"schedule"`

	Args []byte `json:"args"`
}

// UpdateCron: update an existing cron
//
// Update the cron associated with the given id.
func (s *API) UpdateCron(req *UpdateCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CronID) == "" {
		return nil, errors.New("field CronID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteCronRequest struct {
	Region scw.Region `json:"-"`

	CronID string `json:"-"`
}

// DeleteCron: delete an existing cron
//
// Delete the cron associated with the given id.
func (s *API) DeleteCron(req *DeleteCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CronID) == "" {
		return nil, errors.New("field CronID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
		Headers: http.Header{},
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListLogsRequest struct {
	Region scw.Region `json:"-"`

	ContainerID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: timestamp_desc
	OrderBy ListLogsRequestOrderBy `json:"-"`
}

// ListLogs: list your container logs
func (s *API) ListLogs(req *ListLogsRequest, opts ...scw.RequestOption) (*ListLogsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ContainerID) == "" {
		return nil, errors.New("field ContainerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "/logs",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListDomainsRequest struct {
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListDomainsRequestOrderBy `json:"-"`

	ContainerID string `json:"-"`
}

func (s *API) ListDomains(req *ListDomainsRequest, opts ...scw.RequestOption) (*ListDomainsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "container_id", req.ContainerID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetDomainRequest struct {
	Region scw.Region `json:"-"`

	DomainID string `json:"-"`
}

func (s *API) GetDomain(req *GetDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
		Headers: http.Header{},
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateDomainRequest struct {
	Region scw.Region `json:"-"`

	Hostname string `json:"hostname"`

	ContainerID string `json:"container_id"`
}

func (s *API) CreateDomain(req *CreateDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteDomainRequest struct {
	Region scw.Region `json:"-"`

	DomainID string `json:"-"`
}

func (s *API) DeleteDomain(req *DeleteDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
		Headers: http.Header{},
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type IssueJWTRequest struct {
	Region scw.Region `json:"-"`

	ContainerID *string `json:"-"`

	NamespaceID *string `json:"-"`

	ExpiresAt *time.Time `json:"-"`
}

func (s *API) IssueJWT(req *IssueJWTRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "container_id", req.ContainerID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "expires_at", req.ExpiresAt)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/issue-jwt",
		Query:   query,
		Headers: http.Header{},
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListNamespacesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Namespaces = append(r.Namespaces, results.Namespaces...)
	r.TotalCount += uint32(len(results.Namespaces))
	return uint32(len(results.Namespaces)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListContainersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListContainersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListContainersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Containers = append(r.Containers, results.Containers...)
	r.TotalCount += uint32(len(results.Containers))
	return uint32(len(results.Containers)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCronsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCronsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListCronsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Crons = append(r.Crons, results.Crons...)
	r.TotalCount += uint32(len(results.Crons))
	return uint32(len(results.Crons)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLogsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLogsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListLogsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Logs = append(r.Logs, results.Logs...)
	r.TotalCount += uint32(len(results.Logs))
	return uint32(len(results.Logs)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDomainsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Domains = append(r.Domains, results.Domains...)
	r.TotalCount += uint32(len(results.Domains))
	return uint32(len(results.Domains)), nil
}
