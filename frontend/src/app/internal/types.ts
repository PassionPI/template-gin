import type {
  Action,
  AnyAction,
  Dispatch,
  Middleware,
  Reducer,
  Slice,
  ThunkAction,
  ThunkDispatch,
} from "@reduxjs/toolkit";

export type BaseSlice = Record<string, Slice>;

export type GetAppReducer<AppSlice extends BaseSlice> = {
  [K in keyof AppSlice]: AppSlice[K]["reducer"];
};

export type GetAppActions<AppSlice extends BaseSlice> = {
  [K in keyof AppSlice]: AppSlice[K]["actions"];
};

export type GetAppState<AppSlice extends BaseSlice> = {
  [K in keyof GetAppReducer<AppSlice>]: GetAppReducer<AppSlice>[K] extends Reducer<
    infer State
  >
    ? State
    : never;
};

export type GetAppDispatch<AppSlice extends BaseSlice> = ThunkDispatch<
  GetAppState<AppSlice>,
  undefined,
  AnyAction
> &
  Dispatch<AnyAction>;

export type GetAppAction<AppSlice extends BaseSlice> = Parameters<
  GetAppDispatch<AppSlice>
>[0];

export type GetAppMiddleware<AppSlice extends BaseSlice> = Middleware<
  any,
  GetAppState<AppSlice>,
  GetAppDispatch<AppSlice>
>;

export type GetAppThunk<
  AppSlice extends BaseSlice,
  ReturnType = void
> = ThunkAction<ReturnType, GetAppState<AppSlice>, unknown, Action<string>>;

export type GetAppAsyncThunkAPI<
  AppSlice extends BaseSlice,
  Config = {
    /** type of the `extra` argument for the thunk middleware, which will be passed in as `thunkApi.extra` */
    extra?: unknown;
    /** type to be passed into `rejectWithValue`'s first argument that will end up on `rejectedAction.payload` */
    rejectValue?: unknown;
    /** return type of the `serializeError` option callback */
    serializedErrorType?: unknown;
    /** type to be returned from the `getPendingMeta` option callback & merged into `pendingAction.meta` */
    pendingMeta?: unknown;
    /** type to be passed into the second argument of `fulfillWithValue` to finally be merged into `fulfilledAction.meta` */
    fulfilledMeta?: unknown;
    /** type to be passed into the second argument of `rejectWithValue` to finally be merged into `rejectedAction.meta` */
    rejectedMeta?: unknown;
  }
> = {
  state: GetAppState<AppSlice>;
  dispatch: GetAppDispatch<AppSlice>;
} & Config;
