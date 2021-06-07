(self.webpackChunkvachil_booking_svcs_with_prime_ng=self.webpackChunkvachil_booking_svcs_with_prime_ng||[]).push([[611],{1611:(t,e,i)=>{"use strict";i.r(e),i.d(e,{BookingModule:()=>x});var o=i(1455),n=i(8619);let a=(()=>{class t{}return t.\u0275fac=function(e){return new(e||t)},t.\u0275cmp=n.Xpm({type:t,selectors:[["ng-component"]],decls:3,vars:0,consts:[[1,"p-4"],[1,"container"]],template:function(t,e){1&t&&(n.TgZ(0,"div",0),n.TgZ(1,"div",1),n._UZ(2,"router-outlet"),n.qZA(),n.qZA())},directives:[o.lC],encapsulation:2}),t})();var l=i(5108),s=i(2290),r=i(8974),c=i(7894),g=i(322),u=i(1709),p=i(1462),d=i(9892),h=i(6726),k=i(6239),b=i(1116);function D(t,e){if(1&t){const t=n.EpF();n.TgZ(0,"div",10),n.TgZ(1,"h5",11),n._uU(2,"Manage Your Bookings"),n.qZA(),n.TgZ(3,"span",12),n._UZ(4,"i",13),n.TgZ(5,"input",14),n.NdJ("input",function(e){return n.CHM(t),n.oxw(),n.MAs(2).filterGlobal(e.target.value,"contains")}),n.qZA(),n.qZA(),n.qZA()}}function v(t,e){1&t&&(n.TgZ(0,"tr"),n.TgZ(1,"th"),n._uU(2,"Start Date"),n.qZA(),n.TgZ(3,"th"),n._uU(4,"End Date"),n.qZA(),n.TgZ(5,"th"),n._uU(6,"Image"),n.qZA(),n.TgZ(7,"th"),n._uU(8,"Price"),n.qZA(),n._UZ(9,"th"),n.qZA())}function Z(t,e){if(1&t){const t=n.EpF();n.TgZ(0,"tr"),n.TgZ(1,"td"),n._uU(2),n.ALo(3,"date"),n.qZA(),n.TgZ(4,"td"),n._uU(5),n.ALo(6,"date"),n.qZA(),n.TgZ(7,"td"),n._UZ(8,"img",15),n.qZA(),n.TgZ(9,"td"),n._uU(10),n.ALo(11,"currency"),n.qZA(),n.TgZ(12,"td"),n.TgZ(13,"button",16),n.NdJ("click",function(){const e=n.CHM(t).$implicit;return n.oxw().editBooking(e)}),n.qZA(),n.qZA(),n.qZA()}if(2&t){const t=e.$implicit,i=n.oxw();n.xp6(2),n.Oqu(n.lcZ(3,6,t.startDate)),n.xp6(3),n.Oqu(n.lcZ(6,8,t.endDate)),n.xp6(3),n.Q6J("src",i.getBookingVachilTotalCost(t.vachilID),n.LSH)("alt",t.vachilID),n.xp6(2),n.Oqu(n.xi3(11,10,i.getBookingPrice(t.totalPriceID),"INR")),n.xp6(3),n.Q6J("disabled",i.isActive(t.endDate))}}function m(t,e){if(1&t&&(n.TgZ(0,"div",10),n._uU(1),n.qZA()),2&t){const t=n.oxw();n.xp6(1),n.hij(" In total there are ",t.allBooking?t.allBooking.length:0," Bookings. ")}}function B(t,e){if(1&t){const t=n.EpF();n.TgZ(0,"button",17),n.NdJ("click",function(){return n.CHM(t),n.oxw().hideDialog()}),n.qZA(),n.TgZ(1,"button",18),n.NdJ("click",function(){return n.CHM(t),n.oxw().saveBooking()}),n.qZA()}}const f=function(){return["startDate","endDate","totalPriceID"]},A=function(){return{"margin-left":"80px","margin-right":"80px"}};class w{constructor(){this.allBookedVachil=[],this.allBookedPrice=[]}}class T{}const q=[{path:"",component:a,children:[{path:"",component:(()=>{class t{constructor(t,e,i,o,n){this.productService=t,this.messageService=e,this.confirmationService=i,this.acountSvcs=o,this.updateBookingSvcs=n}ngOnInit(){this.productService.fetch({id:this.acountSvcs.currentUser.id}).subscribe(t=>{this.allBooking=t.data.allBookingWithId,this.bookingVachiIdAndPrice=this.getAllBookingDetail(),this.getAllDetail()},t=>console.error(t))}getAllDetail(){this.bookingDetailFull=new T,this.acountSvcs.bookingDetail(this.bookingVachiIdAndPrice).subscribe(t=>{this.bookingDetailFull.allBookedPrice=t.allBookedPrice?new Map(t.allBookedPrice.map(t=>[t.id,t.price])):new Map,this.bookingDetailFull.allBookedVachil=t.allBookedVachil?new Map(t.allBookedVachil.map(t=>[t.id,t.images])):new Map,console.log(this.bookingDetailFull)})}getAllBookingDetail(){let t=new w;for(let e=0;e<this.allBooking.length;e++){let i=this.allBooking[e].totalPriceID,o=this.allBooking[e].vachilID;i&&t.allBookedPrice.push(i),o&&t.allBookedVachil.push(o)}return t}openNew(){this.booking={startDate:new Date,endDate:new Date,id:0,totalPriceID:0,userID:0,vachilID:0},this.submitted=!1,this.bookingDilog=!0}editBooking(t){this.rangeDates=[],this.booking=Object.assign({},t),this.rangeDates.push(new Date(this.booking.startDate)),this.rangeDates.push(new Date(this.booking.endDate)),this.bookingDilog=!0}hideDialog(){this.bookingDilog=!1,this.submitted=!1}saveBooking(){this.submitted=!0,this.booking.startDate=this.rangeDates[0],this.booking.endDate=this.rangeDates[1],this.booking.id&&(this.updateBookingSvcs.mutate({bookingID:this.booking.id,startDate:this.booking.startDate,endDate:this.booking.endDate,userID:this.acountSvcs.currentUser.id,vachilID:this.booking.vachilID}).subscribe(t=>{var e;console.log(null===(e=t.data)||void 0===e?void 0:e.updateBooking)}),this.messageService.add({severity:"success",summary:"Successful",detail:"Booking Created",life:3e3})),this.bookingDilog=!1}findIndexById(t){let e=-1;for(let i=0;i<this.allBooking.length;i++)if(this.allBooking[i].id==t){e=i;break}return e}getBookingVachilTotalCost(t){var e;return null===(e=this.bookingDetailFull.allBookedVachil)||void 0===e?void 0:e.get(t)}getBookingPrice(t){var e,i;return(null===(e=this.bookingDetailFull.allBookedPrice)||void 0===e?void 0:e.get(t))?null===(i=this.bookingDetailFull.allBookedPrice)||void 0===i?void 0:i.get(t):0}isActive(t){let e=new Date;return new Date(t)<e}}return t.\u0275fac=function(e){return new(e||t)(n.Y36(l.E),n.Y36(s.ez),n.Y36(s.YP),n.Y36(r.BR),n.Y36(l.QC))},t.\u0275cmp=n.Xpm({type:t,selectors:[["ng-component"]],decls:13,vars:15,consts:[[1,"card"],["dataKey","id","currentPageReportTemplate","Showing {first} to {last} of {totalRecords} entries",3,"value","rows","paginator","globalFilterFields","rowHover","showCurrentPageReport"],["dt",""],["pTemplate","caption"],["pTemplate","header"],["pTemplate","body"],["pTemplate","summary"],["header","Book vachil","styleClass","p-fluid",3,"visible","modal","visibleChange"],["selectionMode","range","inputId","range",3,"ngModel","readonlyInput","showTime","ngModelChange"],["pTemplate","footer"],[1,"p-d-flex","p-ai-center","p-jc-between"],[1,"p-m-0"],[1,"p-input-icon-left"],[1,"pi","pi-search"],["pInputText","","type","text","placeholder","Search...",3,"input"],["width","100",1,"p-shadow-4",3,"src","alt"],["pButton","","pRipple","","icon","pi pi-pencil",1,"p-button-rounded","p-button-success","p-mr-2",3,"disabled","click"],["pButton","","pRipple","","label","Cancel","icon","pi pi-times",1,"p-button-text",3,"click"],["pButton","","pRipple","","label","Save","icon","pi pi-check",1,"p-button-text",3,"click"]],template:function(t,e){1&t&&(n.TgZ(0,"div",0),n.TgZ(1,"p-table",1,2),n.YNc(3,D,6,0,"ng-template",3),n.YNc(4,v,10,0,"ng-template",4),n.YNc(5,Z,14,13,"ng-template",5),n.YNc(6,m,2,1,"ng-template",6),n.qZA(),n.qZA(),n.TgZ(7,"p-dialog",7),n.NdJ("visibleChange",function(t){return e.bookingDilog=t}),n.TgZ(8,"p-calendar",8),n.NdJ("ngModelChange",function(t){return e.rangeDates=t}),n.qZA(),n._UZ(9,"br"),n._UZ(10,"br"),n._UZ(11,"br"),n.YNc(12,B,2,0,"ng-template",9),n.qZA()),2&t&&(n.xp6(1),n.Q6J("value",e.allBooking)("rows",5)("paginator",!0)("globalFilterFields",n.DdM(13,f))("rowHover",!0)("showCurrentPageReport",!0),n.xp6(6),n.Akn(n.DdM(14,A)),n.Q6J("visible",e.bookingDilog)("modal",!0),n.xp6(1),n.Q6J("ngModel",e.rangeDates)("readonlyInput",!1)("showTime",!0))},directives:[c.iA,s.jx,g.V,u.f,p.JJ,p.On,d.o,h.Hq,k.H],pipes:[b.uU,b.H9],encapsulation:2}),t})()}]}];let I=(()=>{class t{}return t.\u0275fac=function(e){return new(e||t)},t.\u0275mod=n.oAB({type:t}),t.\u0275inj=n.cJS({imports:[[o.Bz.forChild(q)],o.Bz]}),t})();var _=i(5902);let x=(()=>{class t{}return t.\u0275fac=function(e){return new(e||t)},t.\u0275mod=n.oAB({type:t}),t.\u0275inj=n.cJS({imports:[[I,_.m]]}),t})()}}]);