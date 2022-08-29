package models

import "fmt"

type Request struct {
	Operation string `json:"method"`
	Parameters interface{} `json:"parameters"`
}

type DefaultOperationMetadata struct {
	Strategy string `json:"strategy,omitempty"`
	RuningTime string  `json:"time_complexity,omitempty"`
	Space string `json:"space_complexity,omitempty"`
}


func NewDefaultOperationMetadata() *DefaultOperationMetadata {
	m := &DefaultOperationMetadata{}
	return m
}

func (m *DefaultOperationMetadata) SetStrategy(strategy string) {
	m.Strategy = strategy
}

func (m *DefaultOperationMetadata) SetRuningTime(timeComplexity string) {
	m.Strategy = timeComplexity
}

func (m *DefaultOperationMetadata) SetSpace(spaceComplexity string) {
	m.Strategy = spaceComplexity
}

func NewRequest(operation string, params interface{}) *Request {
	r := &Request{}
	r.Operation = operation
	r.Parameters = params
	return r
}
func (r *Request) Update(operation string, params interface{}) {
	r.Operation = operation
	r.Parameters = params
}

type Response struct {
	RequestData Request `json:"requestdata,omitempty"`
	OperationsMetadata interface{} `json:"operation_metadata,omitempty"`
	Result interface{} `json:"result"`
	Errors []string `json:"errors,omitempty"`
}

func NewResponse() *Response {
	response := &Response{}
	response.Errors = make([]string, 0, 8)
	return response
}

func (r *Response) AddError(err error) {
	r.Errors = append(r.Errors, err.Error())
}
func (r *Response) SetResult(result interface{}) {
	r.Result=result
}

func (r *Response) SetOperationsMetadata(meta interface{}) {
	r.OperationsMetadata = meta
}


type OperationInformation struct {
	Operation string `json:"method"`
	Strategy interface {} `json:"strategy"`
	StrategyThresholds interface {} `json:"strategy_bands"`
}
type EmptyParametersError struct {
	message string
}

type CastParametersError struct {
	message string
}

func NewEmptyParametersError(operation string, numberOfParameters int ) *EmptyParametersError {
	epE := &EmptyParametersError{}
	epE.message = fmt.Sprintf("operation %s expects at least %d paramters",operation,numberOfParameters)
	return epE
}
func (epE *EmptyParametersError) Error() string {
	return epE.message
}

func NewCastParametersError(operation string, typePram string ) *CastParametersError {
	cpE := &CastParametersError{}
	cpE.message = fmt.Sprintf("operation %s expects parameters of type %s",operation,typePram)
	return cpE
}
func (epE *CastParametersError) Error() string {
	return epE.message
}
