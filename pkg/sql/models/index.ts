// Code generated by tygo. DO NOT EDIT.

//////////
// source: contact.go

export interface Contact {
  created_at: string /* RFC3339 */;
  email: string;
  id: number /* int64 */;
  info: Info;
  name: string;
  phone: string;
  updated_at: string /* RFC3339 */;
}
export interface Info {
  age: number /* int */;
}
