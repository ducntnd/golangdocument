package main

import (
	"github.com/kataras/iris/v12"
)


func main()  {
	app := iris.New()

	app.HandleDir("/css", iris.Dir("./css"))

	app.RegisterView(iris.HTML("./html",".html").Reload(true))

	app.Get("/", func(ctx iris.Context) {
		//if err := ctx.View("Currency.html"); err != nil {
		//	ctx.StopWithError(iris.StatusInternalServerError, err)
		//	return
		//}
		ctx.ViewData("countries",readCSVCurrencyRate())
		currencyA:=ctx.Values().Get("currencyA")
		currencyB:=ctx.Values().Get("currencyB")
		amount:=ctx.Values().Get("amount")
		checkPost:=ctx.Values().Get("checkPost")
		exchangeA:=ctx.Values().Get("exchangeA")
		exchangeB:=ctx.Values().Get("exchangeB")
		result:=ctx.Values().Get("result")
		ctx.ViewData("currencyA",currencyA)
		ctx.ViewData("currencyB",currencyB)
		ctx.ViewData("amount",amount)
		ctx.ViewData("checkPost",checkPost)
		ctx.ViewData("exchangeA",exchangeA)
		ctx.ViewData("exchangeB",exchangeB)
		ctx.ViewData("result",result)
		ctx.View("Currency")
	})

	app.Post("/", func(ctx iris.Context) {
		moneyresult := MoneyConvertResult{}
		err := ctx.ReadForm(&moneyresult)
		if err != nil {
			if !iris.IsErrPath(err) /* see: https://github.com/kataras/iris/issues/1157 */ ||
				err == iris.ErrEmptyForm {
				ctx.StopWithError(iris.StatusInternalServerError, err)
				return
			}
		}
		ctx.Values().Set("currencyA",moneyresult.CurrencyA)
		ctx.Values().Set("currencyB",moneyresult.CurrencyB)
		ctx.Values().Set("amount",moneyresult.Amount)
		ctx.Values().Set("result",getResult(moneyresult.Amount,moneyresult.CurrencyA,moneyresult.CurrencyB))
		ctx.Values().Set("checkPost",true)
		ctx.Values().Set("exchangeA",getResult(1,moneyresult.CurrencyA,moneyresult.CurrencyB))
		ctx.Values().Set("exchangeB",getResult(1,moneyresult.CurrencyB,moneyresult.CurrencyA))
		ctx.Exec("GET","/")
	})

	app.Listen(":8080")
}
