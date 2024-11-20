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

// AIService  defines AI Service.
type AIService struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired behavior of this APIPlan.
	Spec AIServiceSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen=true

// AIServiceSpec the plugin configuration.
type AIServiceSpec struct {
	Anthropic   *Anthropic   `json:"anthropic,omitempty"`
	AzureOpenAI *AzureOpenAi `json:"azureOpenAi,omitempty"`
	Bedrock     *Bedrock     `json:"bedrock,omitempty"`
	Cohere      *Cohere      `json:"cohere,omitempty"`
	Gemini      *Gemini      `json:"gemini,omitempty"`
	Mistral     *Mistral     `json:"mistral,omitempty"`
	Ollama      *Ollama      `json:"ollama,omitempty"`
	OpenAI      *OpenAi      `json:"openAi,omitempty"`
}

// +k8s:deepcopy-gen=true

// Anthropic configures anthropic.
type Anthropic struct {
	Token  string  `json:"token,omitempty"`
	Model  string  `json:"model,omitempty"`
	Params *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// AzureOpenAi configures AzureOpenAi.
type AzureOpenAi struct {
	APIKey         string  `json:"apiKey,omitempty"`
	Model          string  `json:"model,omitempty"`
	DeploymentName string  `json:"deploymentName,omitempty"`
	BaseURL        string  `json:"baseUrl,omitempty"`
	Params         *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// Bedrock configures Bedrock.
type Bedrock struct {
	Model         string  `json:"model,omitempty"`
	Region        string  `json:"region,omitempty"`
	SystemMessage bool    `json:"systemMessage,string,omitempty"`
	Params        *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// Cohere configures Cohere.
type Cohere struct {
	Token  string  `json:"token,omitempty"`
	Model  string  `json:"model,omitempty"`
	Params *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// Gemini configures Gemini.
type Gemini struct {
	APIKey string  `json:"apiKey,omitempty"`
	Model  string  `json:"model,omitempty"`
	Params *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// Mistral configures Mistral.
type Mistral struct {
	APIKey string  `json:"apiKey,omitempty"`
	Model  string  `json:"model,omitempty"`
	Params *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// Ollama configures Ollama.
type Ollama struct {
	Model   string  `json:"model,omitempty"`
	BaseURL string  `json:"baseUrl,omitempty"`
	Params  *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// OpenAi configures OpenAi.
type OpenAi struct {
	Token  string  `json:"token,omitempty"`
	Model  string  `json:"model,omitempty"`
	Params *Params `json:"params,omitempty"`
}

// +k8s:deepcopy-gen=true

// Params holds LLM params.
type Params struct {
	Temperature      float32 `json:"temperature,string,omitempty"`
	TopP             float32 `json:"topP,string,omitempty"`
	MaxTokens        int     `json:"maxTokens,string,omitempty"`
	FrequencyPenalty float32 `json:"frequencyPenalty,string,omitempty"`
	PresencePenalty  float32 `json:"presencePenalty,string,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AIServiceList defines a list of AIService.
type AIServiceList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AIService `json:"items"`
}
