import type { Definition, Visibility } from "@/types";
import DefinitionView from "@/components/definition-view";
import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Button } from "@/components/ui/button";
import { ArrowUp } from "lucide-react";
import { Separator } from "@/components/ui/separator";
import { useMutation } from "@tanstack/react-query";
import LoadSwap from "@/components/custom-ui/load-swap";
import api from "@/api";
import { useState } from "react";
import { toast } from "sonner";
import ProfilePicture from "@/components/custom-ui/profile-picture";
import { useUser } from "@/context/user";
import { Card, CardContent } from "@/components/ui/card";
import { queryClient } from "@/lib/query";
import { useNavigate } from "@tanstack/react-router";

type PostFormProps = {
  definition: Definition;
};

export default function PostForm({ definition }: PostFormProps) {
  const user = useUser();
  const navigate = useNavigate();
  const [visibility, setVisibility] = useState<Visibility>("public");
  const { mutate, isPending } = useMutation({
    mutationKey: ["post"],
    mutationFn: async () => api.createPost({ content: definition, visibility }),
    onError: (error) => toast.error(error.message, { richColors: true }),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["posts", user.id] });
      navigate({ to: "/app/profile" });
    },
  });

  return (
    <div className="flex-1 place-items-center grid">
      <Card className="max-w-lg">
        <CardContent className="space-y-4">
          <div className="flex gap-4">
            <ProfilePicture
              username={user.username}
              color={user.color}
              imageUrl={user.imageUrl}
            />
            <div>
              <p className="text-muted-foreground text-sm">@{user.username}</p>
              <p>{user.displayname}</p>
            </div>
          </div>
          <DefinitionView definition={definition} />
          <Tabs
            className="w-full"
            value={visibility}
            onValueChange={(v) => setVisibility(v as Visibility)}
          >
            <TabsList className="w-full">
              <TabsTrigger value="public">Public</TabsTrigger>
              <TabsTrigger value="private">Private</TabsTrigger>
            </TabsList>
          </Tabs>
          <Separator />
          <Button
            className="w-full"
            disabled={isPending}
            onClick={() => mutate()}
          >
            <LoadSwap isLoading={isPending} label="Posting...">
              <ArrowUp />
              <span>Post</span>
            </LoadSwap>
          </Button>
        </CardContent>
      </Card>
    </div>
  );
}
