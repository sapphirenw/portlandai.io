package routes

import (
	"fmt"
	"net/http"

	"github.com/sapphirenw/gotmpl"
	"github.com/sapphirenw/portlandai.io/src/components"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// create content
	obj := make(map[string]any, 0)
	obj["Offerings"] = []*components.Card1{
		{
			Title:       "AI Tools",
			Description: "Find a neat AI Tool? Portland AI will vet your choice and work with you to implement your vision.",
			Src:         "/public/images/ai-tools.jpg",
			Alt:         "AI Tools",
			Route:       "/ai-tools",
		},
		{
			Title:       "Custom Integrations",
			Description: "We are experts in utilizing the powerful APIs from OpenAI, Athropic, and more to build truly custom solutions.",
			Src:         "/public/images/custom-software.jpg",
			Alt:         "Custom Software",
			Route:       "/custom-software",
		},
		{
			Title:       "Consultations",
			Description: "We are not interested on selling you on a modern hype train. Meet with us to decide if AI is right for you.",
			Src:         "/public/images/software-consultant.jpeg",
			Alt:         "Software Consultants",
			Route:       "/contact",
		},
	}
	obj["Portfolio"] = []*components.Card1{
		{
			Title:       "Fitness AI",
			Description: "Currently in the works, we are taking our personal fitness tracking app and supercharging it with AI. This will be the best fitness AI on the market when released.",
			Src:         "/public/images/workoutnotepad.co.png",
			Alt:         "Workout Notepad App",
			Route:       "https://workoutnotepad.co",
			External:    true,
			Annotation:  "Comming Soon",
		},
		{
			Title:       "Automatic Blog",
			Description: "We developed a novel system to develop well-researched articles covering a wide array of topics. This product seamlessly combines AI with traditional programming paradigms.",
			Src:         "/public/images/blog.workoutnotepad.co.png",
			Alt:         "Workout Notepad Blog",
			Route:       "https://blog.workoutnotepad.co",
			External:    true,
		},
		{
			Title:       "Multi-Model Chat Bots",
			Description: "Having a little fun, we uses the APIs from mutltiple LLM providers to build a hockey player chat bot with different personalities. Find tuned prompts make for interesting answers.",
			Src:         "/public/images/chuckbot.png",
			Alt:         "ChuckBot",
			Route:       "https://pucknorris.com/chat",
			External:    true,
		},
		{
			Title:       "Custom AI Architecture",
			Description: "As an academic exercise in 2022, one of our engineers developed a neural network (a primitive AI architectural component) from scratch as an exercise of AI competency. Before the days of LLMs.",
			Src:         "/public/images/nn.jakelanders.com.png",
			Alt:         "Neural Network",
			Route:       "https://nn.jakelanders.com",
			External:    true,
		},
		{
			Title:       "Low-Level Convolutions",
			Description: "As another academic exercise, one of our engineers implemented an image kerneling algorithm from scratch. Convultions are common operations found in convolutional neural networks.",
			Src:         "/public/images/convolutions.png",
			Alt:         "Image Convolutions",
			Route:       "https://www.youtube.com/watch?v=P4i2HhJoxeI",
			External:    true,
		},
	}
	obj["Technologies"] = []*components.Technology1{
		{Title: "OpenAI", Route: "https://openai.com"},
		{Title: "Anthropic", Route: "https://anthropic.com"},
		{Title: "Google Gemini", Route: "https://deepmind.google/technologies/gemini/", Img: "https://wp.technologyreview.com/wp-content/uploads/2023/12/gemini_mm_03.png?fit=2160,1214"},
		{Title: "Amazon Bedrock", Route: "https://aws.amazon.com/bedrock/", Img: "https://www.outsystems.com/Forge_CW/_image.aspx/Q8LvY--6WakOw9afDCuuGQ_Q2qNoQaT-xrNXdmgM4dI=/aws-bedrock-connector-2023-01-04%2000-00-00-2023-12-05%2016-01-19"},
		{Title: "Azure ML", Route: "https://azure.microsoft.com/en-us/products/machine-learning/", Img: "https://logowik.com/content/uploads/images/azure-machine-learning-service1395.jpg"},
		{Title: "Bing", Route: "https://www.bing.com"},
		{Title: "Midjourney", Route: "https://www.midjourney.com", Img: "https://asset.brandfetch.io/idPKhKpL-q/id6WXQL8CB.png"},
		{Title: "Hugging Face", Route: "https://huggingface.co"},
		{Title: "LLama", Route: "https://github.com/ggerganov/llama.cpp", Img: "https://camo.githubusercontent.com/b1a52a9f33eedc5ae3701ca5fad90e3bb92587fbaa77f7aea988a0df148083fb/68747470733a2f2f706c2d7075626c69632d646174612e73332e616d617a6f6e6177732e636f6d2f6173736574735f6c696768746e696e672f4c69745f4c4c614d415f496c6c757374726174696f6e33782e706e67"},
		{Title: "Wastonx AI", Route: "https://www.ibm.com/products/watsonx-ai", Img: "https://www.techspot.com/images2/news/bigimage/2023/05/2023-05-09-image-7-j.webp"},
		{Title: "Jasper AI", Route: "https://www.jasper.ai", Img: "https://asset.brandfetch.io/idbrKrUZrF/idB58Dl7XP.svg"},
		{Title: "Ocoya", Route: "https://www.ocoya.com", Img: "https://asset.brandfetch.io/id1rysWj4w/ideRIQSGwk.svg"},
		{Title: "ClickUp AI", Route: "https://clickup.com"},
		{Title: "Copy AI", Route: "https://www.copy.ai", Img: "https://asset.brandfetch.io/idqxaIdXya/idJy6kquFH.png"},
		{Title: "Goose AI", Route: "https://goose.ai"},
		{Title: "Cloudfare AI", Route: "https://developers.cloudflare.com/workers-ai/models/text-generation/", Img: "https://cdn.icon-icons.com/icons2/2699/PNG/512/cloudflare_logo_icon_170372.png"},
	}

	if err := gotmpl.XT.ExecuteTemplate(w, "index.html", obj); err != nil {
		fmt.Println(err)
		Error(w, r)
	}
}
