import{p as e,D as t,s,o as a,c as i,w as n,n as r,i as o,b as l,g as c,t as d,z as u}from"./index-0cdf8aad.js";import{r as f}from"./request.4f472fa9.js";import{_ as p}from"./_plugin-vue_export-helper.1b428a4d.js";const g=""+new URL("editIcon-0d5e9a95.svg",import.meta.url).href;const m=p({data:()=>({app:e(),topHeight:0,userInfo:{},nickName:"",avatar:"",phone:"",statisticsList:[{title:"运动时长",num:0},{title:"累计天数",num:0},{title:"连续天数",num:0}],reservationInfo:[],serviceList:[{icon:"/static/images/mine/orderIcon.svg",title:"我的订单"},{icon:"/static/images/mine/orderIcon.svg",title:"我的预约"},{icon:"/static/images/mine/setIcon.svg",title:"设置"},{icon:"/static/images/mine/agreementIcon.svg",title:"意见反馈"},{icon:"/static/images/mine/contactIcon.svg",title:"联系客服"}],num:"",con:{siteNo:"",startTime:"",endTime:"",title:"",icon:""}}),async mounted(){let t=await e().getUserInfo("reGet");this.userInfo=t,this.$nextTick((()=>this.ready())),this.statisticsList[0].num=t.total_length,this.statisticsList[1].num=t.sport_day,this.statisticsList[2].num=t.total_count,this.serviceList.pop()},methods:{ready(){this.getTopHeight(),this.getRecentlyReserve()},dealWithDate(e){let t=["周日","周一","周二","周三","周四","周五","周六"][new Date(e).getDay()],s="",a=e.split("-");return s=a[0]==(new Date).getFullYear()?a[1]+"-"+a[2]:a[0]+"-"+a[1]+"-"+a[2],{date:s,day:t}},handleContact(e){console.log(e.detail.path),console.log(e.detail.query)},async getRecentlyReserve(){let e=await this.app.getEnum();f({url:"wx/recently/reserve",method:"POST",data:{user_ouid:this.userInfo.ouid}}).then((t=>{if(!t.data)return!1;let s=t.data,a=s.data.slice(0,10),i=this.dealWithDate(a),n=[],r=[],o=s.site_detail;o?(o.forEach((t=>{t.time_enum.forEach((s=>{var a,i;r.push({siteNo:t.site_name,startTime:null==(a=e[s])?void 0:a.split("~")[0],endTime:null==(i=e[s])?void 0:i.split("~")[1]})}))})),n.push({gymnasiumName:s.shop_name,date:i.date,day:i.day,order_no:s.order_no,siteList:r})):n=[],this.reservationInfo=n}))},toDetail(){t({url:"/pages/reservationInfo/reservationInfo?order_no="+this.reservationInfo[0].order_no})},getTopHeight(){let e=s();this.topHeight=e.statusBarHeight+30},toInfoEdit(){t({url:"/pages/infoEdit/infoEdit",events:{updateInfo:async()=>{let t=await e().getUserInfo("reGet");this.userInfo=t}}})},toOrderList(){t({url:"/pages/orderList/orderList"})},toReservationList(){t({url:"/pages/reservationList/reservationList"})},toFeedBack(){t({url:"/pages/feedback/feedback"})},toAboutUs(){t({url:"/pages/aboutUs/aboutUs"})},chooseServiceItem(e){let s="";switch(e.currentTarget.dataset.item){case 0:s="/pages/orderList/orderList";break;case 1:s="/pages/reservationList/reservationList";break;case 2:s="/pages/settings/settings";break;case 3:s="/pages/feedback/feedback";break;case 4:uni.openCustomerServiceChat({extInfo:{url:""},corpId:"",success(e){}})}t({url:s})}}},[["render",function(e,t,s,f,p,m){const h=u,v=o;return a(),i(v,{class:"page",style:r("padding-top: "+p.topHeight+"px;")},{default:n((()=>[l(v,{class:"basicTop flex align-center justify-between"},{default:n((()=>[l(v,{class:"flex align-center"},{default:n((()=>[l(h,{src:p.userInfo.avatar,style:{width:"128rpx",height:"128rpx","border-radius":"50%"},mode:""},null,8,["src"]),l(v,{class:"flex flex-direction align-start",style:{"margin-left":"26rpx"},onClick:m.toInfoEdit},{default:n((()=>[l(v,{class:"nickName"},{default:n((()=>[c(d(p.userInfo.name),1)])),_:1}),l(v,{class:"phone"},{default:n((()=>[c(d(p.userInfo.phone?p.userInfo.phone:""),1)])),_:1})])),_:1},8,["onClick"])])),_:1}),l(v,{onClick:m.toInfoEdit},{default:n((()=>[l(h,{src:g,style:{width:"40rpx",height:"40rpx"},mode:""})])),_:1},8,["onClick"])])),_:1}),l(v,{class:"myReservation",style:{"margin-top":"50rpx"}},{default:n((()=>[l(v,{class:"flex align-center justify-between",onClick:m.toOrderList},{default:n((()=>[l(v,{class:"myReservationTitle"},{default:n((()=>[c("我的订单")])),_:1}),l(v,null,{default:n((()=>[l(h,{src:g,style:{width:"32rpx",height:"32rpx"},mode:""})])),_:1})])),_:1},8,["onClick"])])),_:1}),l(v,{class:"myReservation"},{default:n((()=>[l(v,{class:"flex align-center justify-between",onClick:m.toReservationList},{default:n((()=>[l(v,{class:"myReservationTitle"},{default:n((()=>[c("我的预约")])),_:1}),l(v,null,{default:n((()=>[l(h,{src:g,style:{width:"32rpx",height:"32rpx"},mode:""})])),_:1})])),_:1},8,["onClick"])])),_:1}),l(v,{class:"myReservation"},{default:n((()=>[l(v,{class:"flex align-center justify-between",onClick:m.toFeedBack},{default:n((()=>[l(v,{class:"myReservationTitle"},{default:n((()=>[c("意见反馈")])),_:1}),l(v,null,{default:n((()=>[l(h,{src:g,style:{width:"32rpx",height:"32rpx"},mode:""})])),_:1})])),_:1},8,["onClick"])])),_:1}),l(v,{class:"myReservation"},{default:n((()=>[l(v,{class:"flex align-center justify-between",onClick:m.toAboutUs},{default:n((()=>[l(v,{class:"myReservationTitle"},{default:n((()=>[c("关于我们")])),_:1}),l(v,null,{default:n((()=>[l(h,{src:g,style:{width:"32rpx",height:"32rpx"},mode:""})])),_:1})])),_:1},8,["onClick"])])),_:1})])),_:1},8,["style"])}],["__scopeId","data-v-3f361220"]]);export{m as default};