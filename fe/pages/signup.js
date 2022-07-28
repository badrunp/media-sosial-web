import { useState } from "react";
import Label from "@/components/form/Label";
import Input from "@/components/form/Input";
import Button from "@/components/form/Button";
import Link from "next/link";
import SigninWithEmail from "@/components/form/SigninWithEmail";

export default function Signin() {
  const [data, setData] = useState({
    username: "",
    email: "",
    password: "",
  });

  const handleChangeInput = (e) => {
    const name = e.target.name;
    const value = e.target.value;

    setData({ ...data, [name]: value });
  };

  const handleSignup = async () => {
    return;
    try {
      //   const res = await fetch("http://localhost:8000/api/auth/register", {
      //     method: "POST",
      //     headers: {
      //       "Content-Type": "application/json",
      //     },
      //     body: JSON.stringify(data),
      //   });
      //   console.log(await res.json());
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div className="w-full min-h-screen flex-center">
      <div className="signin w-full max-w-xs sm:max-w-md py-10">
        <div>
          <h1 className="font-bold text-3xl">Sign up</h1>
        </div>

        <form className="mt-6 sm:mt-8">
          <div>
            <Label name="username" label="Username" />
            <Input
              type="text"
              name="username"
              placeholder="john_doe"
              onChange={handleChangeInput}
            />
          </div>

          <div className="mt-4 sm:mt-5">
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

          <div className="mt-4 sm:mt-5">
            <Button handleClick={handleSignup}>Sign up</Button>
          </div>
        </form>

        <SigninWithEmail />

        <div className="mt-8">
          <p className="font-semibold text-sm sm:text-base">
            You have registered?{" "}
            <Link href={"/signin"}>
              <a className="text-violet-600">Sign in</a>
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
}

Signin.guest = true;
