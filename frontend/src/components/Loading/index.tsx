import { LoadingOutlined } from "@ant-design/icons";
import type { SpinProps } from "antd";
import { Spin } from "antd";
import type { FC } from "react";
import { memo } from "react";
import styles from "./index.module.css";

export const BaseLoading = memo(({ className }: { className?: string }) => (
  <div className={className ?? styles.spin}>
    <LoadingOutlined />
  </div>
));

const Loading: FC<Omit<SpinProps, "indicator">> = memo((props) => (
  <Spin {...props} indicator={<BaseLoading className={styles.pure} />} />
));

Loading.defaultProps = {};

export default Loading;
