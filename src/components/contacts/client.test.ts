import createClient from "openapi-fetch";
import type { paths } from "~/src/schema";
import { expect, test } from "bun:test";

test("Create", async () => {
  const client = createClient<paths>({ baseUrl: "http://localhost:21234" });
  const res = await client.POST("/api/contacts", {
    body: {
      email: "",
      info: {
        age: 0,
      },
      name: "",
      phone: "",
    },
  });

  expect(res.error).toBeUndefined();
  expect(res.data).toEqual({
    $schema: "http://localhost:21234/schemas/Contact.json",
    created_at: res.data!.created_at,
    email: "",
    id: 1,
    info: {
      age: 0,
    },
    name: "",
    phone: "",
    updated_at: res.data!.updated_at,
  });
});
