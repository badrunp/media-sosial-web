import { signIn } from "next-auth/react";
import { useState } from "react";
import Link from "next/link";
import Image from "next/image";

export default function Signin() {
  const [data, setData] = useState({
    email: "",
    password: "",
  });

  const handleSignin = async () => {
    try {
      await signIn("credentials", {
        ...data,
        redirect: false,
      });
    } catch (error) {
      console.log(error);
    } finally {
      return;
    }
  };

  const handleChangeInput = (e) => {
    const name = e.target.name;
    const value = e.target.value;

    setData({ ...data, [name]: value });
  };

  return (
    <div className="w-full h-screen flex justify-center items-center">
      <div className="signin w-full max-w-xs sm:max-w-md">
        <div>
          <h1 className="font-bold text-3xl">Login</h1>
        </div>

        <form className="mt-6 sm:mt-8">
          <div>
            <label
              htmlFor="email"
              className="block font-semibold text-sm sm:text-base"
            >
              Email*
            </label>
            <input
              type="text"
              name="email"
              id="email"
              className="tracking-wide border border-zinc-400/70 rounded-full w-full focus:outline-none px-6 py-3 mt-1 placeholder:font-light placeholder:text-zinc-500 text-sm sm:text-base focus:ring-2 focus:border-violet-600 focus:ring-violet-200 transition-colors ease-in-out"
              autoComplete="off"
              placeholder="your@gmail.com"
              onChange={handleChangeInput}
            />
          </div>
          <div className="mt-4 sm:mt-5">
            <label
              htmlFor="password"
              className="block font-semibold text-sm sm:text-base"
            >
              Password*
            </label>
            <input
              type="text"
              name="password"
              id="password"
              className="tracking-wide border border-zinc-400/70 rounded-full w-full focus:outline-none px-6 py-3 mt-1 placeholder:font-light placeholder:text-zinc-500 text-sm sm:text-base focus:ring-2 focus:border-violet-600 focus:ring-violet-200 transition-colors ease-in-out"
              autoComplete="off"
              placeholder="Min. 6 character"
              onChange={handleChangeInput}
            />
          </div>

          <div className="flex justify-between items-center mt-4 sm:mt-5">
            <div className="flex items-center space-x-2">
              <input
                type="checkbox"
                className="rounded-sm border-zinc-400/70 cursor-pointer text-violet-600 focus:ring-violet-600 block font-semibold text-sm sm:text-base"
                id="remember"
              />
              <label
                className="block font-semibold text-sm sm:text-base cursor-pointer"
                htmlFor="remember"
              >
                Remember me
              </label>
            </div>

            <Link href={"/"}>
              <a className="text-violet-600 font-semibold text-sm sm:text-base">
                Forget Password?
              </a>
            </Link>
          </div>

          <button
            type="button"
            className="w-full px-6 py-3 mt-4 sm:mt-5 bg-violet-600 rounded-full text-white font-semibold text-sm sm:text-base"
            onClick={handleSignin}
          >
            Login
          </button>
        </form>

        <div className="mt-8 border-t border-zinc-400/70 pt-8 relative">
          <span className="absolute -top-3 text-sm left-1/2 -translate-x-1/2 bg-white px-4 font-normal text-zinc-500 tracking-wide truncate">
            or Sign in with Email
          </span>
          <button
            type="button"
            className="w-full space-x-3 px-6 py-3 border border-zinc-400/70 rounded-full font-semibold text-sm sm:text-base flex items-center justify-center"
          >
            <Image src={"/images/google.png"} width="25" height="25" />
            <span>Sign in with Google</span>
          </button>
        </div>
      </div>
    </div>
  );
}

Signin.guest = true;
