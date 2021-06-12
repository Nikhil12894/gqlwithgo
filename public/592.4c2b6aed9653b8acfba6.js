(self.webpackChunkvachil_booking_svcs_with_prime_ng=self.webpackChunkvachil_booking_svcs_with_prime_ng||[]).push([[592],{5108:(t,e,r)=>{"use strict";r.d(e,{BK:()=>c,Tb:()=>u,gj:()=>l,$5:()=>p,E:()=>I,WM:()=>D,QC:()=>v,T$:()=>y,ds:()=>b,AT:()=>_,t5:()=>N,h4:()=>F});var n=r(3802),a=r(1783),o=r(8619);const i=n.ZP`
  query allvachil {
    vachil {
      id
      type
      brand
      regNo
      name
      capacity
      unitPrice
      images
    }
  }
`;let c=(()=>{class t extends a.AE{constructor(t){super(t),this.document=i}}return t.\u0275fac=function(e){return new(e||t)(o.LFG(a._M))},t.\u0275prov=o.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();const s=n.ZP`
  mutation saveVachil(
    $type: String!
    $name: String!
    $brand: String!
    $regNo: String!
    $capacity: Int!
    $unitPrice: Float!
    $images: String!
  ) {
    createVachil(
      input: {
        type: $type
        brand: $brand
        name: $name
        regNo: $regNo
        capacity: $capacity
        unitPrice: $unitPrice
        images: $images
      }
    ) {
      id
      type
      brand
      regNo
      name
      capacity
      unitPrice
      images
    }
  }
`;let u=(()=>{class t extends a.mm{constructor(t){super(t),this.document=s}}return t.\u0275fac=function(e){return new(e||t)(o.LFG(a._M))},t.\u0275prov=o.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();const d=n.ZP`
  mutation updateVachils(
    $id: Int!
    $type: String!
    $name: String!
    $brand: String!
    $regNo: String!
    $capacity: Int!
    $unitPrice: Float!
    $images: String!
  ) {
    updateVachil(
      vachilId: $id
      input: {
        type: $type
        brand: $brand
        name: $name
        regNo: $regNo
        capacity: $capacity
        unitPrice: $unitPrice
        images: $images
      }
    ) {
      id
      type
      brand
      regNo
      name
      capacity
      unitPrice
      images
    }
  }
`;let l=(()=>{class t extends a.mm{constructor(t){super(t),this.document=d}}return t.\u0275fac=function(e){return new(e||t)(o.LFG(a._M))},t.\u0275prov=o.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();const m=n.ZP`
  mutation deleteVachil($vachilID: Int!) {
    deleteVachil(vachilId: $vachilID)
  }
`;let p=(()=>{class t extends a.mm{constructor(t){super(t),this.document=m}}return t.\u0275fac=function(e){return new(e||t)(o.LFG(a._M))},t.\u0275prov=o.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();const $=n.ZP`
  query allBookingForUser($id: Int!) {
    allBookingWithId(userID: $id) {
      id
      startDate
      endDate
      vachilID
      totalPriceID
    }
  }
`;let I=(()=>{class t extends a.AE{constructor(t){super(t),this.document=$}}return t.\u0275fac=function(e){return new(e||t)(o.LFG(a._M))},t.\u0275prov=o.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();const g=n.ZP`
  mutation addBooking(
    $startDate: Time!
    $endDate: Time!
    $userID: Int!
    $vachilID: Int!
  ) {
    createBooking(
      input: {
        startDate: $startDate
        endDate: $endDate
        userID: $userID
        vachilID: $vachilID
      }
    ) {
      id
      startDate
      endDate
      userID
      vachilID
      totalPriceID
    }
  }
`;let D=(()=>{class t extends a.mm{constructor(t){super(t),this.document=g}}return t.\u0275fac=function(e){return new(e||t)(o.LFG(a._M))},t.\u0275prov=o.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();const f=n.ZP`
  mutation updateBooking(
    $bookingID: Int!
    $startDate: Time!
    $endDate: Time!
    $userID: Int!
    $vachilID: Int!
  ) {
    updateBooking(
      bookingId: $bookingID
      input: {
        startDate: $startDate
        endDate: $endDate
        userID: $userID
        vachilID: $vachilID
      }
    ) {
      id
      startDate
      endDate
      userID
      vachilID
      totalPriceID
    }
  }
`;let v=(()=>{class t extends a.mm{constructor(t){super(t),this.document=f}}return t.\u0275fac=function(e){return new(e||t)(o.LFG(a._M))},t.\u0275prov=o.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();const h=n.ZP`
  mutation deleteBooking($bookingID: Int!) {
    deleteBooking(bookingId: $bookingID)
  }
`;let y=(()=>{class t extends a.mm{constructor(t){super(t),this.document=h}}return t.\u0275fac=function(e){return new(e||t)(o.LFG(a._M))},t.\u0275prov=o.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();const k=n.ZP`
  query AllBooking($startDate: Time!, $endDate: Time!) {
    allBooking(startDate: $startDate, endDate: $endDate) {
      id
      startDate
      endDate
      userID
      vachilID
      totalPriceID
    }
  }
`;let b=(()=>{class t extends a.AE{constructor(t){super(t),this.document=k}}return t.\u0275fac=function(e){return new(e||t)(o.LFG(a._M))},t.\u0275prov=o.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();const P=n.ZP`
  query AllUsers($userID: Int!) {
    allUsers(userID: $userID) {
      id
      firstName
      lastName
      email
      mobile
      password
      isActive
      userType
      image
    }
  }
`;let _=(()=>{class t extends a.AE{constructor(t){super(t),this.document=P}}return t.\u0275fac=function(e){return new(e||t)(o.LFG(a._M))},t.\u0275prov=o.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();const w=n.ZP`
  mutation deleteUser($userId: Int!) {
    deleteUser(userID: $userId)
  }
`;let N=(()=>{class t extends a.mm{constructor(t){super(t),this.document=w}}return t.\u0275fac=function(e){return new(e||t)(o.LFG(a._M))},t.\u0275prov=o.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();const S=n.ZP`
  mutation updateUser(
    $id: Int!
    $firstName: String!
    $lastName: String!
    $email: String!
    $mobile: String!
    $password: String!
    $image: String!
    $userType: String!
  ) {
    updateUser(
      id: $id
      input: {
        firstName: $firstName
        lastName: $lastName
        email: $email
        mobile: $mobile
        password: $password
        image: $image
        userType: $userType
      }
    ) {
      id
      firstName
      lastName
      email
      mobile
      password
      isActive
      userType
      image
    }
  }
`;let F=(()=>{class t extends a.mm{constructor(t){super(t),this.document=S}}return t.\u0275fac=function(e){return new(e||t)(o.LFG(a._M))},t.\u0275prov=o.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})()}}]);