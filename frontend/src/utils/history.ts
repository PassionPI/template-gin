import { dispatch } from "@/app";
import { slice } from "@/app/slice";
import { SliceHistory } from "@/app/slice/history";

export const nav = (payload: string | SliceHistory) => {
  if (typeof payload === "string") {
    payload = { to: payload };
  }
  dispatch(slice.history.actions.push(payload));
};
