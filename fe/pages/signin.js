import { signIn } from "next-auth/react";
import { useState } from "react";
import Link from "next/link";
import Label from "@/components/form/Label";
import Input from "@/components/form/Input";
import Button from "@/components/form/Button";
import Checkbox from "@/components/form/Checkbox";
import SigninWithEmail from "@/components/form/SigninWithEmail";

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
    <div className="w-full min-h-screen flex-center">
      <div className="signin w-full max-w-xs sm:max-w-md py-10">
        <div>
          <h1 className="font-bold text-3xl">Sign in</h1>
        </div>

        <form className="mt-6 sm:mt-8">
          <div>
            <Label name="email" label="Email" />
            <Input
              type="text"
              name="email"
              placeholder="johndoe@gmail.com"
              onChange={handleChangeInput}
            />
          </div>

          <div className="mt-4 sm:mt-5">
            <Label name="password" label="Password" />
            <Input
              type="password"
              name="password"
              placeholder="Min. 6 character"
              onChange={handleChangeInput}
            />
          </div>

          <div className="flex justify-between items-center mt-4 sm:mt-5">
            <div className="flex items-center space-x-2">
              <Checkbox name="remember" />
              <Label name="remember" label="Remember me" />
            </div>

            <Link href={"/"}>
              <a className="text-violet-600 font-semibold text-sm sm:text-base">
                Forget Password?
              </a>
            </Link>
          </div>

          <div className="mt-4 sm:mt-5">
            <Button handleClick={handleSignin}>Sign in</Button>
          </div>
        </form>

        <SigninWithEmail />

        <div className="mt-8">
          <p className="font-semibold text-sm sm:text-base">
            Not registered yet?{" "}
            <Link href={"/signup"}>
              <a className="text-violet-600">Create an Account</a>
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
}

Signin.guest = true;
