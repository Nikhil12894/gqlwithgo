(self.webpackChunkvachil_booking_svcs_with_prime_ng=self.webpackChunkvachil_booking_svcs_with_prime_ng||[]).push([[592],{5108:(t,e,n)=>{"use strict";n.d(e,{BK:()=>c,Tb:()=>u,gj:()=>$,E:()=>p,WM:()=>g,QC:()=>m});var a=n(3802),i=n(1783),r=n(8619);const o=a.ZP`
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
`;let c=(()=>{class t extends i.AE{constructor(t){super(t),this.document=o}}return t.\u0275fac=function(e){return new(e||t)(r.LFG(i._M))},t.\u0275prov=r.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();const s=a.ZP`
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
`;let u=(()=>{class t extends i.mm{constructor(t){super(t),this.document=s}}return t.\u0275fac=function(e){return new(e||t)(r.LFG(i._M))},t.\u0275prov=r.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();const d=a.ZP`
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
`;let $=(()=>{class t extends i.mm{constructor(t){super(t),this.document=d}}return t.\u0275fac=function(e){return new(e||t)(r.LFG(i._M))},t.\u0275prov=r.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();a.ZP`
  mutation deleteVachil($vachilID: Int!) {
    deleteVachil(vachilId: $vachilID)
  }
`;const l=a.ZP`
  query allBookingForUser($id: Int!) {
    allBookingWithId(userID: $id) {
      id
      startDate
      endDate
      vachilID
      totalPriceID
    }
  }
`;let p=(()=>{class t extends i.AE{constructor(t){super(t),this.document=l}}return t.\u0275fac=function(e){return new(e||t)(r.LFG(i._M))},t.\u0275prov=r.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();const I=a.ZP`
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
`;let g=(()=>{class t extends i.mm{constructor(t){super(t),this.document=I}}return t.\u0275fac=function(e){return new(e||t)(r.LFG(i._M))},t.\u0275prov=r.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();const D=a.ZP`
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
`;let m=(()=>{class t extends i.mm{constructor(t){super(t),this.document=D}}return t.\u0275fac=function(e){return new(e||t)(r.LFG(i._M))},t.\u0275prov=r.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();a.ZP`
  mutation deleteBooking($bookingID: Int!) {
    deleteBooking(bookingId: $bookingID)
  }
`}}]);