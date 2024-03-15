package llms

import (
	"context"
)

// LLM is a Large Language Model.
type LLM interface {
	Call(ctx context.Context, prompt string, options ...CallOption) (string, error)
	Generate(ctx context.Context, prompts []string, options ...CallOption) ([]*Generation, error)
}

// ChatLLM is a LLM that can be used for chatting.
type ChatLLM interface {
	Call(ctx context.Context, messages []ChatMessage, options ...CallOption) (*AIChatMessage, error)
	Generate(ctx context.Context, messages [][]ChatMessage, options ...CallOption) ([]*Generation, error)
}

// Model is an interface multi-modal models implement.
// Note: this is an experimental API.
type Model interface {
	// GenerateContent asks the model to generate content from parts.
	GenerateContent(ctx context.Context, parts []ContentPart, options ...CallOption) (*ContentResponse, error)
}

// Generation is a single generation from a LLM.
type Generation struct {
	// Text is the generated text.
	Text string `json:"text"`
	// Message stores the potentially generated message.
	Message *AIChatMessage `json:"message"`
	// GenerationInfo is the generation info. This can contain vendor-specific information.
	GenerationInfo map[string]any `json:"generation_info"`
	// StopReason is the reason the generation stopped.
	StopReason string `json:"stop_reason"`
}

// LLMResult is the class that contains all relevant information for an LLM Result.
type LLMResult struct {
	Generations [][]*Generation
	LLMOutput   map[string]any
}
