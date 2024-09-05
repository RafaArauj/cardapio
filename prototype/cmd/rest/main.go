package rest

import "context"
import (
'github.com/gin-gonic/gin'
'net/http'
'context'
)

// Exemplo de abstração da camada de serviço
type ServiceInterface interface {
	AdicionarProdutoAoCarrinho(ctx context.Context, idProduto string, quantidade int, idCarrinho string) error
}


func NewControladorRest(servico ServicoInterface) *ControladorRest {
	return &ControladorRest{
		servico: servico,
	}
}

// Exemplo de controlador POO
type ControladorRest struct {
	servico ServicoInterface
}

// gin handler -> func (ctx *gin.Context)

type MessageResponse struct {
	Message string `json:"message"`
}

type AdicionarProdutoAoCarrinhoRequest struct {
	IDProduto string `json:"id_produto"`
	Quantidade int `json:"quantidade"`
}

func (c *ControladorRest) AdicionarProdutoAoCarrinho(ctx *gin.Context) {
	idDoCarrinho := ctx.Param("id")

	var request AdicionarProdutoAoCarrinhoRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, MessageResponse{
			Message: err.Error()
		})
		return
	}

	if err := c.servico.AdicionarProdutoAoCarrinho(ctx, request.IDProduto, request.Quantidade, idDoCarrinho); err != nil {
		ctx.JSON(http.StatusBadRequest, MessageResponse{
			Message: err.Error()
		})
		return
	}

	ctx.JSON(http.StatusOK, MessageResponse{Message: "sucesso"})
}

// Arquivo cmd/rest/main.go
package main

import (
"github.com/xxxx/xxxx/controllers/rest"
"github.com/xxxx/xxxx/services"
"github.com/gin-gonic/gin"
"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	// Adicionar Cross Origin Requested Source - CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	carrinhoService := services.NewCarrinhoService(/*...*/)
	controlador := rest.NewControladorRest(carrinhoService)


	carrinhoGroup := router.Group('/cart')

	// /cart/id-do-carrinho/items
	carrinhoGroup.Post('/:id/items', controlador.AdicionarProdutoAoCarrinho)

	router.Run()
}