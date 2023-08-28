import { slice } from "@/app/slice";
import { defineAppMiddleware } from "@/app/thunk/define";

export const autoLoading = () => {
  const start = /.+\/pending$/;
  const end = /(.+\/)(fulfilled|rejected)$/;
  return defineAppMiddleware(
    ({ dispatch }) =>
      (next) =>
      (action: { type: string }) => {
        const { type } = action;
        if (start.test(type)) {
          dispatch(
            slice.loading.actions.setState((state) => {
              state[type] = true;
            })
          );
        }
        if (end.test(type)) {
          dispatch(
            slice.loading.actions.setState((state) => {
              state[type?.replace?.(end, (_, p) => `${p}pending`)] = false;
            })
          );
        }
        return next(action);
      }
  );
};
