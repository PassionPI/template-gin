import { sign_up } from "@/services/login";
import { Button, Form, Input, Typography } from "antd";
import { useNavigate } from "react-router-dom";

const Sign: React.FC = () => {
  const nav = useNavigate();

  const onFinish = async (values: { username: string; password: string }) => {
    const { username, password } = values;

    const [error] = await sign_up({ username, password });
    console.log("onFinish", error);
    if (!error) {
      nav("/");
    }
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh",
      }}
    >
      <Typography.Title level={2}>Sign</Typography.Title>
      <Form
        name="basic"
        labelCol={{ span: 8 }}
        wrapperCol={{ span: 16 }}
        style={{ maxWidth: 600 }}
        initialValues={{ remember: true }}
        onFinish={onFinish}
        autoComplete="off"
      >
        <Form.Item
          label="Username"
          name="username"
          rules={[{ required: true, message: "Please input your username!" }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          label="Password"
          name="password"
          rules={[{ required: true, message: "Please input your password!" }]}
        >
          <Input.Password />
        </Form.Item>

        <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default Sign;
