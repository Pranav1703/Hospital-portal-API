import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useNavigate } from "react-router-dom"
import { Text } from "lucide-react"
import { useState } from "react"
import axios from "axios"


export function LoginForm({
  className,
  ...props
}: React.ComponentProps<"div">) {

  const navigate = useNavigate()
    const [username,setUsername] = useState<string>("")
    const [password,setPass] = useState<string>("")

  const submit = async()=>{

    try {
      const resp = await axios.post("http://localhost:3000/login",{
        username,
        password
      })

      console.log("login status: ",resp.status)

      navigate("/")
    } catch (error) {
      console.log(error)
    }
  }
    
  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card>
        <CardHeader>
          <CardTitle>Login to your account</CardTitle>
          <CardDescription>
            Enter your Username below to login to your account
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form>
            <div className="flex flex-col gap-6">
              <div className="grid gap-3">
                <Label htmlFor="email">Username</Label>
                <Input
                  id="username"
                  type="username"
                  placeholder="Your Name"
                  required
                  value={username}
                  onChange={(e)=>setUsername(e.target.value)}
                />
              </div>
              <div className="grid gap-3">
                <div className="flex items-center">
                  <Label htmlFor="password">Password</Label>
                </div>
                <Input id="password" type="password" required value={password} onChange={(e)=>setPass(e.target.value)}/>
              </div>
              <div className="flex flex-col gap-3">
                <Button type="submit" className="w-full" onClick={(e)=>{e.preventDefault();submit();}}>
                  Login
                </Button>
              </div>
            </div>
            <div className="mt-4 text-center text-sm">
              Don&apos;t have an account?{" "}
                <button onClick={()=>navigate("/signup")}className="underline underline-offset-4">
                  Sign up
                </button>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}
