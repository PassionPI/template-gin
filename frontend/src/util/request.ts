// fetch的封装, api与fetch对象传参保持一致
// 新增能力
// 	- 参数支持search, 参数格式与URLSearchParams参数相同
//	- 自动 parse json
//  - 自动处理错误

import { getToken } from "../service/login";

type Payload = {
  url: string;
  method?: string;
  body?:
    | null
    | undefined
    | Record<string, unknown>
    | string
    | ReadableStream
    | URLSearchParams
    | Blob
    | FormData
    | BufferSource;
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
      }
    ]
  | [
      error: null,
      value: R,
      meta: {
        payload: Payload;
        response: Response;
      }
    ]
>;

const typeJSON = (ContentType?: string | null) => {
  return String(ContentType).includes("application/json");
};

export const fetcher = async <R>(payload: Payload): Result<R> => {
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
    (err) => [err, null] as [Error, null]
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

export const request = async <R>(payload: Payload): Promise<Result<R>> => {
  return fetcher<R>({
    ...payload,
    headers: {
      "Content-Type": "application/json",
      ...payload.headers,
      Authorization: "Bearer " + getToken(),
    },
  });
};
