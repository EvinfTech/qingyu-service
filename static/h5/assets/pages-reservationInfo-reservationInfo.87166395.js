import{p as e,B as s,s as t,J as a,o as n,c as i,w as o,i as l,b as m,n as r,j as u,g as c,t as f,e as d,x as g,F as h,l as _,r as p,z as y,S as x}from"./index-0cdf8aad.js";import{r as I,_ as w}from"./uni-app.es.b79b333a.js";import{_ as C}from"./u-navbar.c8086c10.js";import{_ as k}from"./u-modal.9ca63db5.js";import{r as T}from"./request.4f472fa9.js";import{_ as B}from"./_plugin-vue_export-helper.1b428a4d.js";import"./u-loading-icon.a67cb4cf.js";import"./u-safe-bottom.196165c8.js";const b=B({data:()=>({app:e(),show:!1,order_no:"",code:"",codeImg:"",scrollViewHeight:0,gymnasiumInfo:{name:"",address:"",img:"",sizteNo:0,hour:0,siteList:[],person:"",phone:"",createTime:"",status:"",siteNum:"",remark:""}}),onLoad(e){this.$nextTick((()=>{this.getCenterHeight()})),e.order_no&&(this.order_no=e.order_no),this.initData()},methods:{initData(){let e=this.app.globalData.enumInfo;T({url:"wx/get/reserve/detail",method:"POST",data:{order_no:this.order_no}}).then((s=>{let t=s.data,a=t.site_detail?t.site_detail:[],n=[],i=0,o=a?a.length:0;this.code=t.check_no,this.codeImg=this.app.globalData.httpUrl+t.check_qr,a.forEach((s=>{s.time_enum.forEach((a=>{n.push({siteName:s.site_name,date:t.reserve_date,startTime:e[a].split("~")[0],endTime:e[a].split("~")[1]}),i+=1}))})),this.gymnasiumInfo.name=t.shop_name,this.gymnasiumInfo.address=t.shop_address,this.gymnasiumInfo.img=this.app.globalData.httpUrl+t.shop_avatar,this.gymnasiumInfo.person=t.user_name,this.gymnasiumInfo.phone=t.user_phone,this.gymnasiumInfo.createTime=t.gmt_creat_order,this.gymnasiumInfo.siteList=n,this.gymnasiumInfo.siteNum=o,this.gymnasiumInfo.hour=i,this.gymnasiumInfo.status=t.status,this.gymnasiumInfo.remark=t.remark}))},toCancel(){this.show=!0},cancel(){this.show=!1},confirm(){T({url:"wx/cancel/order",method:"POST",data:{order_no:this.order_no}}).then((e=>{s({title:"取消成功",icon:"none",duration:2e3,success:()=>{this.show=!1;this.getOpenerEventChannel().emit("toChangeReservationState",this.order_no),setTimeout((()=>{this.initData()}),2e3)}})}))},getCenterHeight(){var e=t().windowHeight;let s=this;a().select(".bottomBox").boundingClientRect((t=>{s.scrollViewHeight=e-44-t.height})).exec()}}},[["render",function(e,s,t,a,T,B){const b=I(p("up-icon"),w),j=I(p("u-navbar"),C),v=y,N=l,D=x,H=I(p("u-modal"),k);return n(),i(N,{class:"page"},{default:o((()=>[m(j,{class:"nav-bar",title:"预约信息",safeAreaInsetTop:!0,autoBack:!1,fixed:!1},{left:o((()=>[m(b,{name:"arrow-left",onClick:T.app.toBack},null,8,["onClick"])])),_:1}),m(D,{"scroll-y":!0,style:r(["height: "+T.scrollViewHeight+"px;",{"padding-top":"14rpx"}])},{default:o((()=>[m(N,{class:"codeBox flex flex-direction align-center justify-center"},{default:o((()=>[m(N,{class:"codeImg",style:r({opacity:"U"==T.gymnasiumInfo.status||"C"==T.gymnasiumInfo.status?"0.2":"1"})},{default:o((()=>[m(v,{src:T.codeImg,class:"w-full h-full",mode:""},null,8,["src"])])),_:1},8,["style"]),m(N,{class:u("U"==T.gymnasiumInfo.status||"C"==T.gymnasiumInfo.status?"disabledText":"codeText")},{default:o((()=>[c(f(T.code),1)])),_:1},8,["class"])])),_:1}),m(N,{class:"gymnasiumBox flex align-center"},{default:o((()=>[m(N,{class:"gymnasiumImg"},{default:o((()=>[m(v,{class:"w-full h-full",src:T.gymnasiumInfo.img,mode:""},null,8,["src"])])),_:1}),m(N,{class:"flex flex-direction align-start"},{default:o((()=>[m(N,{class:"gymnasiumName"},{default:o((()=>[c(f(T.gymnasiumInfo.name),1)])),_:1}),m(N,{class:"gymnasiumText"},{default:o((()=>[c("场馆地址："+f(T.gymnasiumInfo.address),1)])),_:1}),m(N,{class:"gymnasiumText"},{default:o((()=>[c("预约场地："+f(T.gymnasiumInfo.siteNum)+"场("+f(T.gymnasiumInfo.hour)+"小时)",1)])),_:1})])),_:1})])),_:1}),m(N,{class:"siteBox"},{default:o((()=>[m(N,{class:"siteBoxTitle"},{default:o((()=>[c("预约场次")])),_:1}),m(N,{class:"siteListBox flex align-center flex-wrap"},{default:o((()=>[(n(!0),d(h,null,g(T.gymnasiumInfo.siteList,((e,s)=>(n(),i(N,{class:"siteListItem flex flex-direction align-start",key:s},{default:o((()=>[m(N,{class:"siteNo"},{default:o((()=>[c(f(e.siteName),1)])),_:2},1024),m(N,{class:"flex align-center"},{default:o((()=>[m(N,{class:"dateBox"},{default:o((()=>[c(f(e.date),1)])),_:2},1024),m(N,{class:"timeRangeBox"},{default:o((()=>[c(f(e.startTime)+"-"+f(e.endTime),1)])),_:2},1024)])),_:2},1024)])),_:2},1024)))),128))])),_:1})])),_:1}),m(N,{class:"personInfo"},{default:o((()=>[m(N,{style:{"margin-bottom":"18rpx"}},{default:o((()=>[c("预约人："+f(T.gymnasiumInfo.person),1)])),_:1}),m(N,{style:{"margin-bottom":"18rpx"}},{default:o((()=>[c("手机号码："+f(T.gymnasiumInfo.phone),1)])),_:1}),m(N,{style:{"margin-bottom":"18rpx"}},{default:o((()=>[c("备注："+f(T.gymnasiumInfo.remark),1)])),_:1}),m(N,null,{default:o((()=>[c("创建时间："+f(T.gymnasiumInfo.createTime),1)])),_:1})])),_:1})])),_:1},8,["style"]),m(N,{class:"bottomBox flex align-center justify-end"},{default:o((()=>["C"==T.gymnasiumInfo.status?(n(),i(N,{key:0,class:"cancelBtn flex align-center justify-center"},{default:o((()=>[c("已取消")])),_:1})):_("",!0),"U"==T.gymnasiumInfo.status?(n(),i(N,{key:1,class:"cancelBtn flex align-center justify-center"},{default:o((()=>[c("已使用")])),_:1})):_("",!0),"U"!=T.gymnasiumInfo.status&&"C"!=T.gymnasiumInfo.status?(n(),i(N,{key:2,class:"cancelBtn flex align-center justify-center",onClick:B.toCancel},{default:o((()=>[c("取消预约")])),_:1},8,["onClick"])):_("",!0)])),_:1}),m(H,{show:T.show,title:"提示",content:"确定要取消预约吗？",showCancelButton:!0,onConfirm:B.confirm,onCancel:B.cancel},null,8,["show","onConfirm","onCancel"])])),_:1})}],["__scopeId","data-v-9d298399"]]);export{b as default};