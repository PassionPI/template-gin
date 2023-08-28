import { HOME_ROUTE } from "@/routes/router";
import { Button, Result } from "antd";
import type { FC } from "react";
import { memo } from "react";
import { useNavigate } from "react-router-dom";

const NotFound: FC = memo(() => {
  const nav = useNavigate();
  return (
    <Result
      status="403"
      title="403"
      subTitle="Sorry, not authorized to access."
      extra={
        <Button type="primary" onClick={() => nav(HOME_ROUTE)}>
          Back Home
        </Button>
      }
    />
  );
});

NotFound.defaultProps = {};

export default NotFound;
