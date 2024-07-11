import type { FC } from "react";
import { memo, useEffect, useState } from "react";
import { BaseLoading } from ".";
import styles from "./index.module.css";

export interface Props {
  delay?: number;
}

const PageLoading: FC<Props> = memo(({ delay }) => {
  const [waiting, setWaiting] = useState(false);

  useEffect(() => {
    const t = setTimeout(() => setWaiting(true), delay);
    return () => clearTimeout(t);
  }, [delay]);

  return waiting ? null : (
    <div className={styles.loading}>
      <BaseLoading />
    </div>
  );
});

PageLoading.defaultProps = {
  delay: 120,
};

export default PageLoading;
