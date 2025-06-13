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
import { useEffect, useState } from "react"
import axios from "axios"


export function SignUp({
  className,
  ...props
}: React.ComponentProps<"div">) {

  const navigate = useNavigate()
  const [username,setUsername] = useState<string>("")
  const [password,setPass] = useState<string>("")
    const [role,setRole] = useState<string>("")

  const submit = async()=>{
    try {
      const resp = await axios.post("http://localhost:3000/signup",{
        username,
        password,
        role
      })
      console.log("Signup status: ",resp.status)
      navigate("/login")
    } catch (error) {
      console.log(error)
    }
  }
  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card>
        <CardHeader>
          <CardTitle>SignUp</CardTitle>
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
                <Label htmlFor="role">Role</Label>
                <select
                  id="role"
                  name="role"
                  required
                  className="h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm shadow-sm placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2"
                  value={role}
                  onChange={(e)=>setRole(e.target.value)}
                >
                  <option value="">Select role</option>
                  <option value="doctor">Doctor</option>
                  <option value="receptionist">Receptionist</option>
                </select>
              </div>

              <div className="grid gap-3">
                <div className="flex items-center">
                  <Label htmlFor="password">Password</Label>
                </div>
                <Input id="password" type="password" required value={password} onChange={(e)=>setPass(e.target.value)}/>
              </div>
              <div className="flex flex-col gap-3">
                <Button type="submit" className="w-full" onClick={(e)=>{e.preventDefault();submit();}}>
                  SignUp
                </Button>
              </div>
            </div>
            <div className="mt-4 text-center text-sm">
              Already have an account?{" "}
                <button onClick={()=>navigate("/login")}className="underline underline-offset-4">
                  Login
                </button>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}
