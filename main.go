package main

import(
    "fmt"
    "github.com/MTplusWebSystem/GoBotKit/system"
    "github.com/MTplusWebSystem/GoBotKit/botkit"
)

func main() {
    var token string = string(system.Scan("./.token"))
    bot := botkit.BotInit{
        Token: token,
    }
    defaultMessage := func(nome string) string {
        return fmt.Sprintf(`
Olá: %s

Bem-vindo ao nosso bot de vendas de SSH e testes! Como posso ajudá-lo hoje? 🤖✨

Digite /loja para conhecer melhor`, nome)
    }

    const startLayout string = `
🌟 Bem-vindo à nossa loja!

🛍️ Explore nossos produtos e descubra o melhor para você!

1️⃣ Testes: Experimente antes de comprar! Oferecemos testes gratuitos para garantir que você faça a escolha certa.

2️⃣ Planos de vendas:
    - Planos Prontos: Escolha entre uma variedade de planos cuidadosamente elaborados para atender às suas necessidades.
    - Personalizado: Criamos um plano sob medida para você, para garantir que atenda exatamente às suas necessidades.

💬 Não hesite em perguntar aos nossos colaboradores sobre nossos testes, planos de vendas e programa Indique e Ganhe. Estamos aqui para tornar sua experiência de compra simples, fácil e gratificante!

Agradecemos por escolher nossa loja. Esperamos servi-lo em breve! usando /home
`
    for {
        if bot.ReceiveData(){
            go func() {
                bot.Handler("callback_query",func(event string){
                    switch event {
                    case "!VerLoja":
                        layout := map[string]interface{}{
                            "inline_keyboard": [][]map[string]interface{}{
                                {
                                    {"text": "⏰ Testar", "callback_data": "!CreateTest"},
                                },
                                {
                                    {"text": "💬 Informações", "callback_data": "!CreateTest"},
                                },
                                {
                                    {"text": "🚀 Planos", "callback_data": "!CreateTest"},
                                },
                            },
                        }
                        
                        bot.ReplyToPhotoButton("./imgs/home.jpg", layout)
                        
                    }
                })
            }()
            go func() {
                bot.Handler("commands",func(event string){
                    switch event {
                        case "/loja":
                            layout := map[string]interface{}{
                                "inline_keyboard":[][]map[string]interface{}{
                                    {
                                        {"text": "VER LOJA 🛍️", "callback_data": "!VerLoja"},
                                    },
                               },
                            }
                            bot.SendButton(startLayout,layout)
                        case "/home":
                            layout := map[string]interface{}{
                                "inline_keyboard": [][]map[string]interface{}{
                                    {
                                        {"text": "⏰ Testar", "callback_data": "!CreateTest"},
                                    },
                                    {
                                        {"text": "💬 Informações", "callback_data": "!CreateTest"},
                                    },
                                    {
                                        {"text": "🚀 Planos", "callback_data": "!CreateTest"},
                                    },
                                },
                            }
                            
                            bot.ReplyToPhotoButton("./imgs/home.jpg", layout)
                        default:
                            bot.SendMessages(defaultMessage(bot.Username))
                    }
                })
            }()
            go func() {
                bot.Handler("messages", func(event string) {
                    switch event{
                        default:
                            bot.SendMessages(defaultMessage(bot.Username)) 
                    }
                })
            }()
        }
    }
}