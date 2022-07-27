import { useEffect } from "react";
import Layout from "./components/Layout";
import { useStore } from "./_app";

export default function Home() {
  const { state } = useStore();

  useEffect(() => {
    console.log(state);
  }, [state]);
  return (
    <Layout>
      <h1>Home</h1>
    </Layout>
  );
}

Home.auth = true;
