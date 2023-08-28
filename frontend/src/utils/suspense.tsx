import ErrorBoundary from "@/components/ErrorBoundary";
import PageLoading from "@/components/Loading/PageLoading";
import type { FC, PropsWithRef } from "react";
import { Suspense, lazy, memo } from "react";

export function suspense<P>(
  factory: () => Promise<{
    default: FC<P>;
  }>,
  config?: { fallback?: JSX.Element }
) {
  const { fallback = <PageLoading /> } = config ?? {};

  const Lazy = lazy(factory);

  return memo((props: P) => (
    <ErrorBoundary>
      <Suspense fallback={fallback}>
        <Lazy {...(props as JSX.IntrinsicAttributes & PropsWithRef<P>)} />
      </Suspense>
    </ErrorBoundary>
  ));
}
