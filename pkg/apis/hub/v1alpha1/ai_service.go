/*
Copyright (C) 2022-2024 Traefik Labs

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AIService is a Kubernetes-like Service to interact with a text-based LLM provider. It defines the parameters and credentials required to interact with various LLM providers.
type AIService struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired behavior of this AIService.
	Spec AIServiceSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen=true

// AIServiceSpec describes the LLM service provider.
type AIServiceSpec struct {
	Anthropic   *Anthropic   `json:"anthropic,omitempty"`
	AzureOpenAI *AzureOpenAI `json:"azureOpenai,omitempty"`
	Bedrock     *Bedrock     `json:"bedrock,omitempty"`
	Cohere      *Cohere      `json:"cohere,omitempty"`
	Gemini      *Gemini      `json:"gemini,omitempty"`
	Mistral     *Mistral     `json:"mistral,omitempty"`
	Ollama      *Ollama      `json:"ollama,omitempty"`
	OpenAI      *OpenAI      `json:"openai,omitempty"`
}

// +k8s:deepcopy-gen=true

// Anthropic configures Anthropic backend.
type Anthropic struct {
	Token  string  `json:"token"`
	Model  string  `json:"model,omitempty"`
	Params *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// AzureOpenAI configures AzureOpenAI.
type AzureOpenAI struct {
	APIKey         string  `json:"apiKey"`
	Model          string  `json:"model,omitempty"`
	DeploymentName string  `json:"deploymentName"`
	BaseURL        string  `json:"baseUrl"`
	Params         *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// Bedrock configures Bedrock backend.
type Bedrock struct {
	Model         string  `json:"model,omitempty"`
	Region        string  `json:"region,omitempty"`
	SystemMessage bool    `json:"systemMessage,string,omitempty"`
	Params        *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// Cohere configures Cohere backend.
type Cohere struct {
	Token  string  `json:"token"`
	Model  string  `json:"model,omitempty"`
	Params *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// Gemini configures Gemini backend.
type Gemini struct {
	APIKey string  `json:"apiKey"`
	Model  string  `json:"model,omitempty"`
	Params *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// Mistral configures Mistral AI backend.
type Mistral struct {
	APIKey string  `json:"apiKey"`
	Model  string  `json:"model,omitempty"`
	Params *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// Ollama configures Ollama backend.
type Ollama struct {
	Model   string  `json:"model,omitempty"`
	BaseURL string  `json:"baseUrl"`
	Params  *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// OpenAI configures OpenAI.
type OpenAI struct {
	Token  string  `json:"token"`
	Model  string  `json:"model,omitempty"`
	Params *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// Params holds the LLM hyperparameters.
type Params struct {
	Temperature      float32 `json:"temperature,omitempty"`
	TopP             float32 `json:"topP,omitempty"`
	MaxTokens        int     `json:"maxTokens,omitempty"`
	FrequencyPenalty float32 `json:"frequencyPenalty,omitempty"`
	PresencePenalty  float32 `json:"presencePenalty,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AIServiceList defines a list of AIService.
type AIServiceList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AIService `json:"items"`
}
