package ai

type RequestSourceData struct {
	Type      string `json:"type"`
	MediaType string `json:"media_type"`
	Data      string `json:"data"`
}

type RequestContent struct {
	Type   string            `json:"type"`
	Text   string            `json:"text,omitempty"`
	Source RequestSourceData `json:"source,omitzero"`
}

type Tool struct {
	Type  string         `json:"type"`
	ID    string         `json:"id"`
	Name  string         `json:"name"`
	Input map[string]any `json:"input,omitempty"`
}

type Request struct {
	//Default fields
	Model     string           `json:"model"`
	MaxTokens int              `json:"max_tokens"`
	Messages  []RequestMessage `json:"messages"`

	//Optional fields
	System                 string         `json:"system,omitempty"`
	Stream                 bool           `json:"stream,omitempty"`
	Temperature            float64        `json:"temperature,omitempty"`
	TopP                   float64        `json:"top_p,omitempty"`
	TopK                   float64        `json:"top_k,omitempty"`
	StopSequences          []string       `json:"stop_sequences,omitempty"`
	Tools                  []Tool         `json:"tools,omitempty"`
	ToolChoice             any            `json:"tool_choice,omitempty"`
	DisableParallelToolUse bool           `json:"disable_parallel_tool_use,omitempty"`
	Thinking               bool           `json:"thinking,omitempty"`
	Metadata               map[string]any `json:"metadata,omitempty"`
}

type RequestMessage struct {
	//Default fields
	Role    string           `json:"role"`
	Content []RequestContent `json:"content"`
}

type Balance struct {
	UsdBalance         string `json:"usd_balance"`
	NanoBalance        string `json:"nano_balance"`
	NanoDepositAddress string `json:"nanoDepositAddress"`
}

type Response struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Role    string `json:"role"`
	Model   string `json:"model"`
	Content []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"content"`
	StopReason   string `json:"stop_reason"`
	StopSequence any    `json:"stop_sequence"`
	Usage        struct {
		InputTokens              int `json:"input_tokens"`
		OutputTokens             int `json:"output_tokens"`
		CacheCreationInputTokens int `json:"cache_creation_input_tokens"`
		CacheReadInputTokens     int `json:"cache_read_input_tokens"`
		CacheCreation            struct {
			Ephemeral5MInputTokens int `json:"ephemeral_5m_input_tokens"`
			Ephemeral1HInputTokens int `json:"ephemeral_1h_input_tokens"`
		} `json:"cache_creation"`
		ServiceTier string `json:"service_tier"`
	} `json:"usage"`
}
