import { StopOutlined } from "@ant-design/icons";
import styles from "./index.module.css";

const ErrorHint = () => {
  return (
    <div className={styles.err}>
      <div className={styles.icon}>
        <StopOutlined />
      </div>
      <code>Error</code>
    </div>
  );
};

export default ErrorHint;
