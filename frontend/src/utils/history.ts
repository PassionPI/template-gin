import { dispatch } from "@/app";
import { slice } from "@/app/slice";
import { SliceHistory } from "@/app/slice/history";

export const nav = (payload: SliceHistory) => {
  dispatch(slice.history.actions.push(payload));
};
