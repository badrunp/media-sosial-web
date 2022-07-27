import NextAuth from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";

export default NextAuth({
  providers: [
    CredentialsProvider({
      async authorize(credentials) {
        try {
          // const res = await fetch("http://localhost:8000/api/auth/login", {
          //   method: "POST",
          //   headers: {
          //     "Content-Type": "application/json",
          //   },
          //   body: JSON.stringify({
          //     email: credentials.email,
          //     password: credentials.password,
          //   }),
          // });
          // const json = await res.json();
          // if (json.status == 200) {
          //   return json.data;
          // }
          return {
            name: "badrun",
            email: "bbadrunn@gmail.com",
            password: "123456",
          };
        } catch (error) {
          console.log(error);
          return null;
        }
      },
    }),
  ],
  session: {
    strategy: "jwt",
  },
});
