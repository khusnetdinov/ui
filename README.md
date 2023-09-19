### Telegrapha.UI

#### Files structure:
```
.
├── api                         # package api
│   ├── api.go                  # Ui.Api Types & Structures for interacting with Api on low lelel
│   ├── go.mod                  # Module file
│   ├── methods.go              # Methods for calliing Http.Request to Telegram Api
│   ├── params.go               # Params for calliing Http.Request to Telegram Api
│   └── types.go                # Types & Structrures for handling Http.Response from Telegram Api
├── go.mod                      # Module file
├── go.sum                      #
├── helpers.go                  # Helpers
├── main.go                     # package main
│── README.md                   # Documatation
│── LICENSE                     # License
└── .gitignore                  # Git ignored files
```
#### Links:
  - Read official docs in [Telegram Api Documentation](https://core.telegram.org/bots/api)
  - Download last actual collection from [Postman collections repository](https://github.com/sys-001/telegram-bot-api-versions)
  - JavaScript [gist](https://gist.github.com/khusnetdinov/fdc60831db748d3d8ead601fbb637411) for get methods & types for docs
