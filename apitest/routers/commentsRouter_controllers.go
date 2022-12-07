package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["apitest/controllers:AL000003Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:AL000003Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:AL000012Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:AL000012Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:AL000013Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:AL000013Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:AL000014Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:AL000014Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:AL000015Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:AL000015Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:AL000018Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:AL000018Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:AL000054Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:AL000054Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:AL000055Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:AL000055Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:AL000060Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:AL000060Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:AL000061Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:AL000061Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:AL000063Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:AL000063Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:AL000065Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:AL000065Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:BP000001Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:BP000001Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:BP000002Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:BP000002Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:BP000003Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:BP000003Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:BP000004Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:BP000004Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:BP000005Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:BP000005Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:BP000006Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:BP000006Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:BP000007Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:BP000007Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:BP000008Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:BP000008Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:BP000031Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:BP000031Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:BP000032Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:BP000032Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:BP000034Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:BP000034Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:BP000035Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:BP000035Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CM000006Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CM000006Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CM000007Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CM000007Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CM000021Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CM000021Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CM000040Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CM000040Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CM000041Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CM000041Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CM000044Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CM000044Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CM000052Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CM000052Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CU000001Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CU000001Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CU000002Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CU000002Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CU000003Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CU000003Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CU000004Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CU000004Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CU000005Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CU000005Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CU000006Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CU000006Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CU000007Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CU000007Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CU000008Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CU000008Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CU000009Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CU000009Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CU000010Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CU000010Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CU000011Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CU000011Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CU000012Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CU000012Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:CU000061Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:CU000061Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000001Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000001Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000002Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000002Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000003Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000003Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000004Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000004Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000005Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000005Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000006Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000006Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000007Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000007Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000008Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000008Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000009Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000009Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000010Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000010Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000011Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000011Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000012Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000012Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000013Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000013Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000014Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000014Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000015Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000015Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000016Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000016Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000017Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000017Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000018Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000018Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000019Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000019Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000020Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000020Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000021Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000021Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000024Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000024Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000025Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000025Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000026Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000026Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000027Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000027Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000029Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000029Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000030Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000030Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000031Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000031Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000035Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000035Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000036Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000036Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000039Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000039Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000040Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000040Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000041Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000041Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000042Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000042Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000043Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000043Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000044Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000044Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000045Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000045Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000046Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000046Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000047Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000047Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000048Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000048Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000049Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000049Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000050Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000050Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000051Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000051Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000052Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000052Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000053Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000053Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000055Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000055Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000056Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000056Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP0000572Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP0000572Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000057Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000057Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000058Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000058Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000059Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000059Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000060Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000060Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000061Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000061Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000062Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000062Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000063Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000063Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000064Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000064Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000065Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000065Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000066Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000066Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000067Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000067Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000068Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000068Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000070Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000070Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000071Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000071Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000072Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000072Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000077Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000077Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000078Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000078Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000097Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000097Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000098Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000098Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000099Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000099Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000100Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000100Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000102Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000102Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000103Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000103Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000104Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000104Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000105Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000105Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000106Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000106Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000107Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000107Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000108Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000108Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000109Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000109Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000110Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000110Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000111Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000111Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000112Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000112Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000113Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000113Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP000114Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP000114Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010001Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010001Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010002Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010002Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010003Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010003Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010004Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010004Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010005Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010005Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010006Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010006Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010007Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010007Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010008Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010008Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010009Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010009Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010010Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010010Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010011Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010011Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010012Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010012Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010013Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010013Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010014Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010014Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010015Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010015Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010016Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010016Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010017Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010017Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010018Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010018Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010019Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010019Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010020Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010020Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010021Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010021Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010022Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010022Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010023Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010023Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010024Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010024Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010025Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010025Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010026Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010026Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP010027Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP010027Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP013015Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP013015Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP013018Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP013018Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP013028Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP013028Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP020001Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP020001Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP020002Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP020002Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP020003Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP020003Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP020004Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP020004Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP020007Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP020007Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP020009Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP020009Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:FP020010Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:FP020010Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000001Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000001Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000002Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000002Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000006Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000006Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000007Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000007Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000008Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000008Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000009Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000009Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000010Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000010Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000011Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000011Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000012Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000012Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000013Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000013Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000014Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000014Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000015Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000015Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000017Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000017Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000018Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000018Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000022Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000022Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000023Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000023Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000024Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000024Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000025Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000025Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000026Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000026Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000027Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000027Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000028Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000028Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000029Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000029Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000030Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000030Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000031Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000031Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000032Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000032Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000034Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000034Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000035Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000035Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000036Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000036Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000037Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000037Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000038Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000038Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000039Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000039Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000040Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000040Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000041Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000041Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000042Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000042Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000043Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000043Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000044Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000044Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000049Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000049Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000050Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000050Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000052Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000052Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000053Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000053Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000054Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000054Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000055Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000055Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000059Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000059Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000060Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000060Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000061Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000061Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000062Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000062Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000069Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000069Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000070Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000070Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000071Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000071Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000072Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000072Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:IC000073Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:IC000073Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:NT000001Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:NT000001Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:NT000002Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:NT000002Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:NT000003Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:NT000003Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:NT000004Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:NT000004Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:NT000005Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:NT000005Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:NT000006Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:NT000006Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:NT000007Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:NT000007Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:NT000008Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:NT000008Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:NT000009Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:NT000009Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:NT000010Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:NT000010Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:NT000011Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:NT000011Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:NT000012Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:NT000012Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:NT000013Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:NT000013Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000001Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000001Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000002Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000002Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000003Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000003Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000004Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000004Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000005Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000005Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000006Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000006Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000009Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000009Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000010Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000010Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000011Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000011Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000012Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000012Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000013Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000013Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000014Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000014Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000015Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000015Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000016Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000016Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000017Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000017Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000018Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000018Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000019Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000019Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000020Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000020Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000021Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000021Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000044Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000044Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000045Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000045Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000046Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000046Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000047Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000047Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000052Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000052Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000053Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000053Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000054Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000054Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000057Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000057Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000058Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000058Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000068Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000068Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000069Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000069Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000070Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000070Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000078Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000078Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000080Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000080Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000082Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000082Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000084Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000084Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000087Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000087Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000089Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000089Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000096Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000096Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PD000100Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PD000100Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PF000001Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PF000001Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PF000002Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PF000002Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PF000003Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PF000003Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PF000004Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PF000004Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PF000005Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PF000005Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PF000006Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PF000006Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PF000007Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PF000007Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PF000008Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PF000008Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PF000009Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PF000009Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PF000010Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PF000010Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PI000001Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PI000001Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PI000002Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PI000002Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PI000003Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PI000003Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PI000004Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PI000004Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PI000005Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PI000005Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PI000006Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PI000006Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PI000007Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PI000007Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PI000008Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PI000008Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PI000009Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PI000009Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PI000010Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PI000010Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PR000001Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PR000001Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PR000002Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PR000002Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PR000003Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PR000003Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PR000004Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PR000004Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PR000005Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PR000005Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PR000009Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PR000009Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PR000010Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PR000010Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PR000012Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PR000012Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:PR000013Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:PR000013Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000001Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000001Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000002Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000002Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000004Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000004Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000005Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000005Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000006Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000006Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000007Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000007Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000008Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000008Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000009Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000009Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000010Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000010Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000011Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000011Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000012Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000012Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000014Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000014Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000015Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000015Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000017Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000017Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000018Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000018Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000019Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000019Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000020Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000020Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000029Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000029Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000030Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000030Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:SMP0000031Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:SMP0000031Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitest/controllers:Test0001Controller"] = append(beego.GlobalControllerRouter["apitest/controllers:Test0001Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
