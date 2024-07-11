import { defineSliceItem } from "@/app/internal/creator";
import { PayloadAction } from "@reduxjs/toolkit";

export type SliceHistory = {
  to:
    | null
    | string
    | {
        pathname: string;
        search?: string;
        hash?: string;
      };
  replace?: boolean;
};

export const history = defineSliceItem({
  initialState: (): SliceHistory => ({
    to: null,
    replace: undefined,
  }),
  reducers: {
    push(state, { payload }: PayloadAction<SliceHistory>) {
      Object.assign(state, payload);
    },
  },
});
