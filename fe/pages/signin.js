import { signIn, useSession } from "next-auth/react";
import Link from "next/link";
import { useRouter } from "next/router";

export default function Signin() {
  const { data: session } = useSession();
  const router = useRouter();
  const handleSignin = async () => {
    const res = await signIn("credentials", { redirect: false });
    const redirect = router.query.redirect;
    if (redirect) {
      router.push(router.query.redirect);
    }
  };
  return (
    <div>
      {session ? (
        <>
          <div>Anda sudah signin</div>
          <Link href={"/"}>
            <a>Kembali</a>
          </Link>
        </>
      ) : (
        <>
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
          <div>
            <label>Email</label>
            <input type="text" placeholder="Email" />
          </div>
          <div>
            <label>Password</label>
            <input type="text" placeholder="password" />
          </div>
          <button onClick={handleSignin}>Signin</button>
        </>
      )}
    </div>
  );
}
