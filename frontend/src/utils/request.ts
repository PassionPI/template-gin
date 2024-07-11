// fetch的封装, api与fetch对象传参保持一致
// 新增能力
// 	- 参数支持search, 参数格式与URLSearchParams参数相同
//	- 自动 parse json
//  - 自动处理错误

import { getToken } from "@/services/login";
import { message } from "antd";

type Payload = {
  url: string;
  method?: string;
  body?: Record<string, unknown> | RequestInit["body"];
  search?: ConstructorParameters<typeof URLSearchParams>[0];
  headers?: Record<string, string>;
};

type Result<R> = Promise<
  | [
      error: Error,
      value: null,
      meta: {
        payload: Payload;
        response: null;
      },
    ]
  | [
      error: null,
      value: R,
      meta: {
        payload: Payload;
        response: Response;
      },
    ]
>;

const typeJSON = (ContentType?: string | null) => {
  return String(ContentType).includes("application/json");
};

const fetcher = async <R>(payload: Payload): Result<R> => {
  const { method, body, search, headers } = payload ?? {};

  const url = new URL(payload?.url, self.location.origin);

  if (search) {
    url.search = new URLSearchParams(search).toString();
  }

  const [error, response] = await fetch(url.toString(), {
    headers,
    method,
    body:
      body != null && typeJSON(headers?.["Content-Type"])
        ? typeof body === "string"
          ? body
          : JSON.stringify(body)
        : (body as RequestInit["body"]),
  }).then(
    (val) => [null, val] as [null, Response],
    (err) => [err, null] as [Error, null],
  );

  if (error) {
    return [
      error,
      null,
      {
        payload,
        response,
      },
    ];
  }

  const value: R = await (typeJSON(response.headers.get("Content-Type"))
    ? response.json().catch(() => ({}))
    : response.text().catch(() => ""));

  return [
    null,
    value,
    {
      payload,
      response,
    },
  ];
};

export const request = async <R>(payload: Payload): Result<R> => {
  return fetcher<{ data: R; error?: string }>({
    ...payload,
    headers: {
      "Content-Type": "application/json",
      ...payload.headers,
      Authorization: "Bearer " + getToken(),
    },
  }).then(([error, value, meta]) => {
    if (error) {
      message.error(error.message);
      return [error, value, meta];
    }
    if (value?.error) {
      message.error(value?.error);
      return [Error(value?.error), null, meta] as unknown as Result<R>;
    }
    return [error, value.data, meta];
  });
};
