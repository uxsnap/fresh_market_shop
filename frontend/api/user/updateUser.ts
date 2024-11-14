import client from "../client";

type Body = {
  firstName: string;
  lastName: string;
  birthday?: string | null;
  email: string;
};

export const updateUser = (body: Body) => {
  return client.put("/user", {
    ...body,
    birthday: body.birthday ?? undefined,
  });
};

updateUser.queryKey = "updateUser";
