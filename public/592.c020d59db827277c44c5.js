(self.webpackChunkvachil_booking_svcs_with_prime_ng=self.webpackChunkvachil_booking_svcs_with_prime_ng||[]).push([[592],{5108:(t,e,i)=>{"use strict";i.d(e,{BK:()=>o,E:()=>d});var a=i(3802),n=i(1783),r=i(8619);const c=a.ZP`
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
`;let o=(()=>{class t extends n.AE{constructor(t){super(t),this.document=c}}return t.\u0275fac=function(e){return new(e||t)(r.LFG(n._M))},t.\u0275prov=r.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})();a.ZP`
  mutation saveVachil(
    $type: String!
    $name: String!
    $brand: String!
    $regNo: String!
    $capacity: Int!
    $unitPrice: Float!
  ) {
    createVachil(
      input: {
        type: $type
        brand: $brand
        name: $name
        regNo: $regNo
        capacity: $capacity
        unitPrice: $unitPrice
      }
    ) {
      id
      type
      brand
      regNo
      name
      capacity
      unitPrice
    }
  }
`,a.ZP`
  mutation updateVachils(
    $id: Int!
    $type: String!
    $name: String!
    $brand: String!
    $regNo: String!
    $capacity: Int!
    $unitPrice: Float!
    $images: [String!]
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
      }
    ) {
      id
      type
      brand
      regNo
      name
      capacity
      unitPrice
    }
  }
`,a.ZP`
  mutation deleteVachil($vachilID: Int!) {
    deleteVachil(vachilId: $vachilID)
  }
`;const u=a.ZP`
  query allBookingForUser($id: Int!) {
    allBookingWithId(userID: $id) {
      id
      startDate
      endDate
      vachilID
      totalPriceID
    }
  }
`;let d=(()=>{class t extends n.AE{constructor(t){super(t),this.document=u}}return t.\u0275fac=function(e){return new(e||t)(r.LFG(n._M))},t.\u0275prov=r.Yz7({token:t,factory:t.\u0275fac,providedIn:"root"}),t})()}}]);