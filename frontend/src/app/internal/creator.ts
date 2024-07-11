import {
  AnyAction,
  AsyncThunk,
  AsyncThunkOptions,
  AsyncThunkPayloadCreator,
  CaseReducer,
  Draft,
  MiddlewareArray,
  PayloadAction,
  ThunkMiddleware,
  configureStore,
  createAsyncThunk,
  createSelector,
  createSlice,
} from "@reduxjs/toolkit";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";
import type {
  BaseSlice,
  GetAppAsyncThunkAPI,
  GetAppDispatch,
  GetAppMiddleware,
  GetAppReducer,
  GetAppState,
  GetAppThunk,
} from "./types";

type GetReturnType<T extends Record<string, (...args: any[]) => any>> = {
  [K in keyof T]: ReturnType<T[K]>;
};

export const creator = <AppSlice extends BaseSlice>({
  slice,
  middleware = [],
}: {
  slice: AppSlice;
  middleware?: GetAppMiddleware<AppSlice>[];
  defaultMiddlewareOption?: {
    thunk?: boolean;
    immutableCheck?: boolean;
    serializableCheck?: boolean;
  };
}) => {
  type AppSliceName = keyof AppSlice;
  type AppReducer = GetAppReducer<AppSlice>;
  type AppState = GetAppState<AppSlice>;

  const { reducer } = Object.entries(slice).reduce(
    (acc, [name, { reducer }]) => {
      acc.reducer[name as AppSliceName] = reducer;
      return acc;
    },
    { reducer: {} } as {
      reducer: AppReducer;
    }
  );
  const store = configureStore<
    AppState,
    AnyAction,
    MiddlewareArray<[ThunkMiddleware<AppState, AnyAction>]>
  >({
    reducer,
    middleware(getDefaultMiddleware) {
      return getDefaultMiddleware({
        serializableCheck: false,
      }).concat(...(middleware as []));
    },
  });
  return {
    store,
  };
};

export const defineSliceItem = <
  T extends Record<string, unknown>,
  R extends Record<string, CaseReducer<T, PayloadAction<any>>> = Record<
    string,
    CaseReducer<T, PayloadAction<any>>
  >
>({
  initialState,
  reducers,
}: {
  initialState: () => T;
  reducers?: R;
}) => {
  return (name: string) => {
    type Reducers = {
      [K in keyof R]: R[K];
    } & {
      setState(
        state: Draft<T>,
        { payload }: PayloadAction<(state: T) => void>
      ): void;
    };
    return createSlice<T, Reducers, string>({
      name,
      initialState,
      reducers: {
        ...(reducers as any),
        setState(state, { payload }: PayloadAction<(state: T) => void>) {
          if (typeof payload === "function") {
            payload(state as T);
          }
        },
      },
    });
  };
};

export const defineSlices = <
  T extends Record<string, ReturnType<typeof defineSliceItem<any>>>
>(
  sliceCollect: T
) => {
  type Key = keyof T;
  return Object.entries(sliceCollect).reduce((acc, [key, call]) => {
    acc[key as Key] = call(key) as ReturnType<T[Key]>;
    return acc;
  }, {} as GetReturnType<T>);
};

export const defineAppHooks = <AppSlice extends BaseSlice>() => {
  const useAppDispatch = () => useDispatch<GetAppDispatch<AppSlice>>();
  const useAppSelector: TypedUseSelectorHook<GetAppState<AppSlice>> =
    useSelector;
  return {
    useAppDispatch,
    useAppSelector,
  };
};

export const definer = <AppSlice extends BaseSlice>() => {
  type AppThunk<Returned = void> = GetAppThunk<AppSlice, Returned>;
  type AppAsyncThunkAPI = GetAppAsyncThunkAPI<AppSlice>;
  type AppMiddleware = GetAppMiddleware<AppSlice>;
  type AppState = GetAppState<AppSlice>;

  type FunctionArrayReturnType<
    T extends ReadonlyArray<(...args: any[]) => unknown>
  > = {
    [K in keyof T]: ReturnType<T[K]>;
  };

  const defineAppSelector = <
    Selectors extends ReadonlyArray<(state: AppState) => unknown>,
    R
  >(
    selectors: Selectors,
    calc: (...selected: FunctionArrayReturnType<Selectors>) => R
  ) => createSelector(...selectors, calc) as unknown as (state: AppState) => R;

  const defineAppThunk = <R>(
    thunk: (...args: Parameters<AppThunk>) => R
  ): AppThunk<R> => thunk;

  const defineAppMiddleware = (fn: AppMiddleware) => fn;

  const defineAppAsyncThunk = <
    Returned,
    ThunkArg = void,
    ThunkAPI extends AppAsyncThunkAPI = AppAsyncThunkAPI
  >(
    typePrefix: string,
    payloadCreator: AsyncThunkPayloadCreator<Returned, ThunkArg, ThunkAPI>,
    options?: AsyncThunkOptions<ThunkArg, ThunkAPI>
  ): AsyncThunk<Returned, ThunkArg, ThunkAPI> => {
    return createAsyncThunk<Returned, ThunkArg, ThunkAPI>(
      typePrefix,
      payloadCreator as AsyncThunkPayloadCreator<Returned, ThunkArg, ThunkAPI>,
      options
    );
  };

  const defineAppAsyncThunkCollection = <
    T extends Record<
      string,
      AsyncThunkPayloadCreator<any, any, AppAsyncThunkAPI>
    >
  >(
    collection: T
  ) => {
    return (namespace: string) => {
      type Returned = {
        [K in keyof T]: AsyncThunk<
          Awaited<ReturnType<T[K]>>,
          Parameters<T[K]>[0],
          GetAppAsyncThunkAPI<AppSlice, AppAsyncThunkAPI>
        >;
      };
      return Object.entries(collection).reduce<Returned>(
        (acc, [key, action]) => {
          acc[key as keyof T] = defineAppAsyncThunk(
            `${namespace}/${key}`,
            action
          );
          return acc;
        },
        {} as Returned
      );
    };
  };
  const defineAppAsyncThunks = <
    T extends Record<string, ReturnType<typeof defineAppAsyncThunkCollection>>
  >(
    thunkCollect: T
  ) => {
    type Key = keyof T;
    return Object.entries(thunkCollect).reduce((acc, [key, call]) => {
      acc[key as Key] = call(`@@async/${key}`) as ReturnType<T[Key]>;
      return acc;
    }, {} as GetReturnType<T>);
  };

  return {
    defineAppSelector,
    defineAppThunk,
    defineAppMiddleware,
    defineAppAsyncThunkCollection,
    defineAppAsyncThunks,
  };
};
