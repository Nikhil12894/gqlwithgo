(self.webpackChunkvachil_booking_svcs_with_prime_ng=self.webpackChunkvachil_booking_svcs_with_prime_ng||[]).push([[611],{1611:(t,e,i)=>{"use strict";i.r(e),i.d(e,{BookingModule:()=>B});var n=i(1455),o=i(8619);let a=(()=>{class t{}return t.\u0275fac=function(e){return new(e||t)},t.\u0275cmp=o.Xpm({type:t,selectors:[["ng-component"]],decls:3,vars:0,consts:[[1,"p-4"],[1,"container"]],template:function(t,e){1&t&&(o.TgZ(0,"div",0),o.TgZ(1,"div",1),o._UZ(2,"router-outlet"),o.qZA(),o.qZA())},directives:[n.lC],encapsulation:2}),t})();var r=i(5108),s=i(2290),c=i(8974),l=i(7894),p=i(322),u=i(1709),g=i(1462),d=i(9892),h=i(6726),Z=i(6239),m=i(1116);function b(t,e){if(1&t){const t=o.EpF();o.TgZ(0,"div",10),o.TgZ(1,"h5",11),o._uU(2,"Manage Your Bookings"),o.qZA(),o.TgZ(3,"span",12),o._UZ(4,"i",13),o.TgZ(5,"input",14),o.NdJ("input",function(e){return o.CHM(t),o.oxw(),o.MAs(2).filterGlobal(e.target.value,"contains")}),o.qZA(),o.qZA(),o.qZA()}}function f(t,e){1&t&&(o.TgZ(0,"tr"),o.TgZ(1,"th"),o._uU(2,"Start Date"),o.qZA(),o.TgZ(3,"th"),o._uU(4,"End Date"),o.qZA(),o.TgZ(5,"th"),o._uU(6,"Image"),o.qZA(),o.TgZ(7,"th"),o._uU(8,"Price"),o.qZA(),o._UZ(9,"th"),o.qZA())}function k(t,e){if(1&t){const t=o.EpF();o.TgZ(0,"tr"),o.TgZ(1,"td"),o._uU(2),o.ALo(3,"date"),o.qZA(),o.TgZ(4,"td"),o._uU(5),o.ALo(6,"date"),o.qZA(),o._UZ(7,"td"),o.TgZ(8,"td"),o._uU(9),o.ALo(10,"currency"),o.qZA(),o.TgZ(11,"td"),o.TgZ(12,"button",15),o.NdJ("click",function(){const e=o.CHM(t).$implicit;return o.oxw().editBooking(e)}),o.qZA(),o.qZA(),o.qZA()}if(2&t){const t=e.$implicit;o.xp6(2),o.Oqu(o.lcZ(3,3,t.startDate)),o.xp6(3),o.Oqu(o.lcZ(6,5,t.endDate)),o.xp6(4),o.Oqu(o.xi3(10,7,t.totalPriceID,"INR"))}}function D(t,e){if(1&t&&(o.TgZ(0,"div",10),o._uU(1),o.qZA()),2&t){const t=o.oxw();o.xp6(1),o.hij(" In total there are ",t.allBooking?t.allBooking.length:0," Bookings. ")}}function v(t,e){if(1&t){const t=o.EpF();o.TgZ(0,"button",16),o.NdJ("click",function(){return o.CHM(t),o.oxw().hideBookingDialog()}),o.qZA(),o.TgZ(1,"button",17),o.NdJ("click",function(){return o.CHM(t),o.oxw().saveBooking()}),o.qZA()}}const T=function(){return["startDate","endDate","totalPriceID"]},A=function(){return{"margin-left":"80px","margin-right":"80px"}},w=[{path:"",component:a,children:[{path:"",component:(()=>{class t{constructor(t,e,i,n){this.productService=t,this.messageService=e,this.confirmationService=i,this.acountSvcs=n}ngOnInit(){this.productService.fetch({id:this.acountSvcs.currentUser.id}).subscribe(t=>{this.allBooking=t.data.allBookingWithId},t=>console.error(t))}openNew(){this.booking={startDate:new Date,endDate:new Date,id:0,totalPriceID:0,userID:0,vachilID:0},this.submitted=!1,this.bookingDilog=!0}editBooking(t){this.rangeDates=[],this.booking=Object.assign({},t),this.rangeDates.push(new Date(this.booking.startDate)),this.rangeDates.push(new Date(this.booking.endDate)),this.bookingDilog=!0}hideDialog(){this.bookingDilog=!1,this.submitted=!1}saveProduct(){this.submitted=!0,this.allBooking.push(this.booking),this.messageService.add({severity:"success",summary:"Successful",detail:"Product Created",life:3e3}),this.allBooking=[...this.allBooking],this.bookingDilog=!1,this.booking={startDate:new Date,endDate:new Date,id:0,totalPriceID:0,userID:0,vachilID:0}}}return t.\u0275fac=function(e){return new(e||t)(o.Y36(r.E),o.Y36(s.ez),o.Y36(s.YP),o.Y36(c.BR))},t.\u0275cmp=o.Xpm({type:t,selectors:[["ng-component"]],decls:13,vars:15,consts:[[1,"card"],["dataKey","id","currentPageReportTemplate","Showing {first} to {last} of {totalRecords} entries",3,"value","rows","paginator","globalFilterFields","rowHover","showCurrentPageReport"],["dt",""],["pTemplate","caption"],["pTemplate","header"],["pTemplate","body"],["pTemplate","summary"],["header","Book vachil","styleClass","p-fluid",3,"visible","modal","visibleChange"],["selectionMode","range","inputId","range",3,"ngModel","readonlyInput","showTime","ngModelChange"],["pTemplate","footer"],[1,"p-d-flex","p-ai-center","p-jc-between"],[1,"p-m-0"],[1,"p-input-icon-left"],[1,"pi","pi-search"],["pInputText","","type","text","placeholder","Search...",3,"input"],["pButton","","pRipple","","icon","pi pi-pencil",1,"p-button-rounded","p-button-success","p-mr-2",3,"click"],["pButton","","pRipple","","label","Cancel","icon","pi pi-times",1,"p-button-text",3,"click"],["pButton","","pRipple","","label","Save","icon","pi pi-check",1,"p-button-text",3,"click"]],template:function(t,e){1&t&&(o.TgZ(0,"div",0),o.TgZ(1,"p-table",1,2),o.YNc(3,b,6,0,"ng-template",3),o.YNc(4,f,10,0,"ng-template",4),o.YNc(5,k,13,10,"ng-template",5),o.YNc(6,D,2,1,"ng-template",6),o.qZA(),o.qZA(),o.TgZ(7,"p-dialog",7),o.NdJ("visibleChange",function(t){return e.bookingDilog=t}),o.TgZ(8,"p-calendar",8),o.NdJ("ngModelChange",function(t){return e.rangeDates=t}),o.qZA(),o._UZ(9,"br"),o._UZ(10,"br"),o._UZ(11,"br"),o.YNc(12,v,2,0,"ng-template",9),o.qZA()),2&t&&(o.xp6(1),o.Q6J("value",e.allBooking)("rows",5)("paginator",!0)("globalFilterFields",o.DdM(13,T))("rowHover",!0)("showCurrentPageReport",!0),o.xp6(6),o.Akn(o.DdM(14,A)),o.Q6J("visible",e.bookingDilog)("modal",!0),o.xp6(1),o.Q6J("ngModel",e.rangeDates)("readonlyInput",!0)("showTime",!0))},directives:[l.iA,s.jx,p.V,u.f,g.JJ,g.On,d.o,h.Hq,Z.H],pipes:[m.uU,m.H9],encapsulation:2}),t})()}]}];let q=(()=>{class t{}return t.\u0275fac=function(e){return new(e||t)},t.\u0275mod=o.oAB({type:t}),t.\u0275inj=o.cJS({imports:[[n.Bz.forChild(w)],n.Bz]}),t})();var _=i(5902);let B=(()=>{class t{}return t.\u0275fac=function(e){return new(e||t)},t.\u0275mod=o.oAB({type:t}),t.\u0275inj=o.cJS({imports:[[q,_.m]]}),t})()}}]);