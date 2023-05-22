import { request } from "../../util/request";

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
    url: "/pub",
    method: "POST",
  });
  if (error) {
    console.error(error);
  } else {
    localStorage.setItem(publicKey, value.publicKey);
  }
};

export const login = async (username: string, password: string) => {
  const [error, value] = await request<{ token: string }>({
    url: "/login",
    method: "POST",
    body: {
      username,
      password,
    },
  });
  if (error) {
    console.error(error);
  } else {
    localStorage.setItem(token, value.token);
  }
};

export const signup = async (username: string, password: string) => {
  const [error, value] = await request<{ token: string }>({
    url: "/sign",
    method: "POST",
    body: {
      username,
      password,
    },
  });
  if (error) {
    console.error(error);
  } else {
    localStorage.setItem(token, value.token);
  }
};
