import { request } from "@/utils/request";
import { rsaEncrypt } from "@/utils/rsa";

const publicKey = "publicKey";
const token = "token";

export const getPubKey = () => localStorage.getItem(publicKey);

export const getToken = () => localStorage.getItem(token);

export const logout = () => localStorage.removeItem(token);

export const isLogin = () => localStorage.getItem(token) != null;

export const invokePubKey = async () => {
  if (getPubKey()) {
    return;
  }
  const [error, value] = await request<{ publicKey: string }>({
    url: "/api/pub",
    method: "POST",
  });
  if (error || !value.publicKey) {
    console.error(error);
  } else {
    localStorage.setItem(publicKey, value.publicKey);
  }
};

export const login = async (body: {
  username: string;
  password: string;
}): Promise<[Error, null] | [null, { token: string }]> => {
  const pubKey = getPubKey();
  if (!pubKey) {
    return [Error("public key not found"), null];
  }
  const encrypt = await rsaEncrypt(pubKey, body?.password);
  const [error, value] = await request<{ token: string }>({
    url: "/api/login",
    method: "POST",
    body: {
      username: body?.username,
      password: encrypt,
    },
  });
  if (error) {
    console.error(error);
  } else {
    localStorage.setItem(token, value.token);
  }
  return [error, value] as [Error, null] | [null, { token: string }];
};

export const sign_up = async (body: {
  username: string;
  password: string;
}): Promise<[Error, null] | [null, { token: string }]> => {
  const pubKey = getPubKey();
  if (!pubKey) {
    return [Error("public key not found"), null];
  }
  const encrypt = await rsaEncrypt(pubKey, body?.password);
  const [error, value] = await request<{ token: string }>({
    url: "/api/sign",
    method: "POST",
    body: {
      username: body?.username,
      password: encrypt,
    },
  });
  if (error) {
    console.error(error);
  } else {
    localStorage.setItem(token, value.token);
  }
  return [error, value] as [Error, null] | [null, { token: string }];
};
