(self.webpackChunkvachil_booking_svcs_with_prime_ng=self.webpackChunkvachil_booking_svcs_with_prime_ng||[]).push([[850],{7850:(r,e,t)=>{"use strict";t.r(e),t.d(e,{AccountModule:()=>J});var s=t(1462),n=t(1116),i=t(1455),o=t(8619),a=t(8974);function u(r,e){1&r&&(o.TgZ(0,"div"),o._uU(1,"Username is required"),o.qZA())}function d(r,e){if(1&r&&(o.TgZ(0,"div",13),o.YNc(1,u,2,0,"div",14),o.qZA()),2&r){const r=o.oxw();o.xp6(1),o.Q6J("ngIf",r.f.username.errors.required)}}function c(r,e){1&r&&(o.TgZ(0,"div"),o._uU(1,"Password is required"),o.qZA())}function f(r,e){if(1&r&&(o.TgZ(0,"div",13),o.YNc(1,c,2,0,"div",14),o.qZA()),2&r){const r=o.oxw();o.xp6(1),o.Q6J("ngIf",r.f.password.errors.required)}}function m(r,e){1&r&&o._UZ(0,"span",15)}const g=function(r){return{"is-invalid":r}};function l(r,e){1&r&&(o.TgZ(0,"div"),o._uU(1,"First Name is required"),o.qZA())}function p(r,e){if(1&r&&(o.TgZ(0,"div",17),o.YNc(1,l,2,0,"div",18),o.qZA()),2&r){const r=o.oxw();o.xp6(1),o.Q6J("ngIf",r.f.firstName.errors.required)}}function Z(r,e){1&r&&(o.TgZ(0,"div"),o._uU(1,"Last Name is required"),o.qZA())}function v(r,e){if(1&r&&(o.TgZ(0,"div",17),o.YNc(1,Z,2,0,"div",18),o.qZA()),2&r){const r=o.oxw();o.xp6(1),o.Q6J("ngIf",r.f.lastName.errors.required)}}function q(r,e){1&r&&(o.TgZ(0,"div"),o._uU(1,"Username is required"),o.qZA())}function b(r,e){if(1&r&&(o.TgZ(0,"div",17),o.YNc(1,q,2,0,"div",18),o.qZA()),2&r){const r=o.oxw();o.xp6(1),o.Q6J("ngIf",r.f.username.errors.required)}}function h(r,e){1&r&&(o.TgZ(0,"div"),o._uU(1,"Password is required"),o.qZA())}function A(r,e){1&r&&(o.TgZ(0,"div"),o._uU(1,"Password must be at least 6 characters"),o.qZA())}function T(r,e){if(1&r&&(o.TgZ(0,"div",17),o.YNc(1,h,2,0,"div",18),o.YNc(2,A,2,0,"div",18),o.qZA()),2&r){const r=o.oxw();o.xp6(1),o.Q6J("ngIf",r.f.password.errors.required),o.xp6(1),o.Q6J("ngIf",r.f.password.errors.minlength)}}function _(r,e){1&r&&o._UZ(0,"span",19)}const N=function(r){return{"is-invalid":r}},w=[{path:"",component:(()=>{class r{constructor(r,e){this.router=r,this.accountService=e,this.accountService.getToken()&&this.router.navigate(["/"])}}return r.\u0275fac=function(e){return new(e||r)(o.Y36(i.F0),o.Y36(a.BR))},r.\u0275cmp=o.Xpm({type:r,selectors:[["ng-component"]],decls:2,vars:0,consts:[[1,"col-md-6","offset-md-3","mt-5"]],template:function(r,e){1&r&&(o.TgZ(0,"div",0),o._UZ(1,"router-outlet"),o.qZA())},directives:[i.lC],encapsulation:2}),r})(),children:[{path:"login",component:(()=>{class r{constructor(r,e,t,s,n){this.formBuilder=r,this.route=e,this.router=t,this.accountService=s,this.alertService=n,this.loading=!1,this.submitted=!1}ngOnInit(){this.form=this.formBuilder.group({username:["",s.kI.required],password:["",s.kI.required]})}get f(){return this.form.controls}onSubmit(){this.submitted=!0,this.alertService.clear(),this.form.invalid||(this.loading=!0,this.accountService.signIn(this.form.value))}}return r.\u0275fac=function(e){return new(e||r)(o.Y36(s.qu),o.Y36(i.gz),o.Y36(i.F0),o.Y36(a.BR),o.Y36(a.c9))},r.\u0275cmp=o.Xpm({type:r,selectors:[["ng-component"]],decls:21,vars:11,consts:[[1,"card"],[1,"card-header"],[1,"card-body"],[3,"formGroup","ngSubmit"],[1,"form-group"],["for","username"],["type","text","formControlName","username",1,"form-control",3,"ngClass"],["class","invalid-feedback",4,"ngIf"],["for","password"],["type","password","formControlName","password",1,"form-control",3,"ngClass"],[1,"btn","btn-primary",3,"disabled"],["class","spinner-border spinner-border-sm mr-1",4,"ngIf"],["routerLink","../register",1,"btn","btn-link"],[1,"invalid-feedback"],[4,"ngIf"],[1,"spinner-border","spinner-border-sm","mr-1"]],template:function(r,e){1&r&&(o.TgZ(0,"div",0),o.TgZ(1,"h4",1),o._uU(2,"Login"),o.qZA(),o.TgZ(3,"div",2),o.TgZ(4,"form",3),o.NdJ("ngSubmit",function(){return e.onSubmit()}),o.TgZ(5,"div",4),o.TgZ(6,"label",5),o._uU(7,"Username"),o.qZA(),o._UZ(8,"input",6),o.YNc(9,d,2,1,"div",7),o.qZA(),o.TgZ(10,"div",4),o.TgZ(11,"label",8),o._uU(12,"Password"),o.qZA(),o._UZ(13,"input",9),o.YNc(14,f,2,1,"div",7),o.qZA(),o.TgZ(15,"div",4),o.TgZ(16,"button",10),o.YNc(17,m,1,0,"span",11),o._uU(18," Login "),o.qZA(),o.TgZ(19,"a",12),o._uU(20,"Register"),o.qZA(),o.qZA(),o.qZA(),o.qZA(),o.qZA()),2&r&&(o.xp6(4),o.Q6J("formGroup",e.form),o.xp6(4),o.Q6J("ngClass",o.VKq(7,g,e.submitted&&e.f.username.errors)),o.xp6(1),o.Q6J("ngIf",e.submitted&&e.f.username.errors),o.xp6(4),o.Q6J("ngClass",o.VKq(9,g,e.submitted&&e.f.password.errors)),o.xp6(1),o.Q6J("ngIf",e.submitted&&e.f.password.errors),o.xp6(2),o.Q6J("disabled",e.loading),o.xp6(1),o.Q6J("ngIf",e.loading))},directives:[s._Y,s.JL,s.sg,s.Fj,s.JJ,s.u,n.mk,n.O5,i.yS],encapsulation:2}),r})()},{path:"register",component:(()=>{class r{constructor(r,e,t,s,n){this.formBuilder=r,this.route=e,this.router=t,this.accountService=s,this.alertService=n,this.loading=!1,this.submitted=!1}ngOnInit(){this.form=this.formBuilder.group({firstName:["",s.kI.required],lastName:["",s.kI.required],username:["",s.kI.required],password:["",[s.kI.required,s.kI.minLength(6)]]})}get f(){return this.form.controls}onSubmit(){this.accountService.signUp(this.form.value).subscribe(r=>{r.result&&(this.form.reset(),this.router.navigate(["../login"],{relativeTo:this.route}))})}}return r.\u0275fac=function(e){return new(e||r)(o.Y36(s.qu),o.Y36(i.gz),o.Y36(i.F0),o.Y36(a.BR),o.Y36(a.c9))},r.\u0275cmp=o.Xpm({type:r,selectors:[["ng-component"]],decls:31,vars:19,consts:[[1,"card"],[1,"card-header"],[1,"card-body"],[3,"formGroup","ngSubmit"],[1,"form-group"],["for","firstName"],["type","text","formControlName","firstName",1,"form-control",3,"ngClass"],["class","invalid-feedback",4,"ngIf"],["for","lastName"],["type","text","formControlName","lastName",1,"form-control",3,"ngClass"],["for","username"],["type","text","formControlName","username",1,"form-control",3,"ngClass"],["for","password"],["type","password","formControlName","password",1,"form-control",3,"ngClass"],[1,"btn","btn-primary",3,"disabled"],["class","spinner-border spinner-border-sm mr-1",4,"ngIf"],["routerLink","../login",1,"btn","btn-link"],[1,"invalid-feedback"],[4,"ngIf"],[1,"spinner-border","spinner-border-sm","mr-1"]],template:function(r,e){1&r&&(o.TgZ(0,"div",0),o.TgZ(1,"h4",1),o._uU(2,"Register"),o.qZA(),o.TgZ(3,"div",2),o.TgZ(4,"form",3),o.NdJ("ngSubmit",function(){return e.onSubmit()}),o.TgZ(5,"div",4),o.TgZ(6,"label",5),o._uU(7,"First Name"),o.qZA(),o._UZ(8,"input",6),o.YNc(9,p,2,1,"div",7),o.qZA(),o.TgZ(10,"div",4),o.TgZ(11,"label",8),o._uU(12,"Last Name"),o.qZA(),o._UZ(13,"input",9),o.YNc(14,v,2,1,"div",7),o.qZA(),o.TgZ(15,"div",4),o.TgZ(16,"label",10),o._uU(17,"Username"),o.qZA(),o._UZ(18,"input",11),o.YNc(19,b,2,1,"div",7),o.qZA(),o.TgZ(20,"div",4),o.TgZ(21,"label",12),o._uU(22,"Password"),o.qZA(),o._UZ(23,"input",13),o.YNc(24,T,3,2,"div",7),o.qZA(),o.TgZ(25,"div",4),o.TgZ(26,"button",14),o.YNc(27,_,1,0,"span",15),o._uU(28," Register "),o.qZA(),o.TgZ(29,"a",16),o._uU(30,"Cancel"),o.qZA(),o.qZA(),o.qZA(),o.qZA(),o.qZA()),2&r&&(o.xp6(4),o.Q6J("formGroup",e.form),o.xp6(4),o.Q6J("ngClass",o.VKq(11,N,e.submitted&&e.f.firstName.errors)),o.xp6(1),o.Q6J("ngIf",e.submitted&&e.f.firstName.errors),o.xp6(4),o.Q6J("ngClass",o.VKq(13,N,e.submitted&&e.f.lastName.errors)),o.xp6(1),o.Q6J("ngIf",e.submitted&&e.f.lastName.errors),o.xp6(4),o.Q6J("ngClass",o.VKq(15,N,e.submitted&&e.f.username.errors)),o.xp6(1),o.Q6J("ngIf",e.submitted&&e.f.username.errors),o.xp6(4),o.Q6J("ngClass",o.VKq(17,N,e.submitted&&e.f.password.errors)),o.xp6(1),o.Q6J("ngIf",e.submitted&&e.f.password.errors),o.xp6(2),o.Q6J("disabled",e.loading),o.xp6(1),o.Q6J("ngIf",e.loading))},directives:[s._Y,s.JL,s.sg,s.Fj,s.JJ,s.u,n.mk,n.O5,i.yS],encapsulation:2}),r})()}]}];let x=(()=>{class r{}return r.\u0275fac=function(e){return new(e||r)},r.\u0275mod=o.oAB({type:r}),r.\u0275inj=o.cJS({imports:[[i.Bz.forChild(w)],i.Bz]}),r})(),J=(()=>{class r{}return r.\u0275fac=function(e){return new(e||r)},r.\u0275mod=o.oAB({type:r}),r.\u0275inj=o.cJS({imports:[[n.ez,s.UX,x]]}),r})()}}]);