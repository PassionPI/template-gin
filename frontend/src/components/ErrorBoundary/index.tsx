import { PureComponent } from "react";
import ErrorHint from "./ErrorHint";

class ErrorBoundary extends PureComponent<{ children: JSX.Element }> {
  state = { hasError: false };

  static getDerivedStateFromError() {
    return { hasError: true };
  }

  componentDidCatch(e: Error) {
    window.console.error("ErrorBoundary", e);
  }

  render() {
    if (this.state.hasError) {
      return <ErrorHint />;
    }
    return this.props.children;
  }
}

export default ErrorBoundary;
