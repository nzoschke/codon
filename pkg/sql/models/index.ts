// Code generated by tygo. DO NOT EDIT.

//////////
// source: contact.go

export interface Contact {
  created_at: string /* RFC3339 */;
  email: string;
  id: number /* int64 */;
  meta: Meta;
  name: string;
  phone: string;
  updated_at: string /* RFC3339 */;
}
export interface Meta {
  age: number /* int */;
}
