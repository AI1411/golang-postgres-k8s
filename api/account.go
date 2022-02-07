package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD,EUR,JPY"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	log.Printf("id %+v", req.ID)
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) createAccount(ctx *gin.Context) {

}

func (server *Server) listAccount(ctx *gin.Context) {

}
