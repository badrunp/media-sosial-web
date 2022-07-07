import { useSession, signIn, signOut } from "next-auth/react";
import { useEffect } from "react";
import Link from "next/link";

export default function Home() {
  const { data: session, status } = useSession({
    required: true,
    onUnauthenticated() {
      console.log("kesini");
    },
  });

  useEffect(() => {
    console.log(status);
  }, [status]);

  useEffect(() => {
    console.log(session);
  }, [session]);

  return (
    <div>
      <h1>Hello World</h1>
      <Link href={"/"}>
        <a>Home</a>
      </Link>
      <br />
      <Link href={"/about"}>
        <a>About</a>
      </Link>
      <br />
      <Link href={"/signin"}>
        <a>Signin</a>
      </Link>
      <br />
      {session ? (
        <button onClick={() => signOut({ redirect: false })}>Keluar</button>
      ) : (
        <button onClick={() => signIn("credentials")}>Masuk</button>
      )}
    </div>
  );
}
