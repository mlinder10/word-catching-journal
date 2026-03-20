import type { Post } from "@/types";
import {
  Card,
  CardAction,
  CardContent,
  CardFooter,
  CardHeader,
} from "../ui/card";
import DefinitionView, { DefinitionViewSkeleton } from "../definition-view";
import { Button } from "../ui/button";
import { Expand } from "lucide-react";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTrigger,
} from "../ui/dialog";
import { Link } from "@tanstack/react-router";
import ProfilePicture from "../custom-ui/profile-picture";
import { PostBookmarkButton, PostLikeButton } from "./post-action-button";

type PostViewProps = {
  post: Post;
};

export function PostView({ post }: PostViewProps) {
  return (
    <Dialog>
      <Card>
        <CardHeader>
          <Link
            to="/app/profile/$profileId"
            params={{ profileId: post.userId }}
            className="flex items-center gap-2"
          >
            <ProfilePicture
              username={post.username}
              imageUrl={post.imageUrl}
              color={post.color}
              size={32}
            />
            <div>
              <p className="text-muted-foreground text-sm">{`@${post.username}`}</p>
              <p>{post.displayname}</p>
            </div>
          </Link>
          <CardAction>
            <DialogTrigger asChild>
              <Button variant="outline">
                <Expand />
              </Button>
            </DialogTrigger>
          </CardAction>
        </CardHeader>
        <CardContent>
          <DefinitionView definition={post} full={false} />
        </CardContent>
        <CardFooter className="flex py-2">
          <PostLikeButton
            postId={post.id}
            count={post.likeCount}
            isMarked={post.isLiked}
          />
          <PostBookmarkButton
            postId={post.id}
            count={post.bookmarkCount}
            isMarked={post.isBookmarked}
          />
          <span className="ml-auto text-muted-foreground text-xs">
            {new Date(post.createdAt).toLocaleDateString()}
          </span>
        </CardFooter>
      </Card>

      <DialogContent>
        <DialogHeader className="flex flex-row items-center gap-2">
          <ProfilePicture
            username={post.username}
            imageUrl={post.imageUrl}
            color={post.color}
          />
          <div>
            <p className="text-muted-foreground text-sm">{`@${post.username}`}</p>
            <p>{post.displayname}</p>
          </div>
        </DialogHeader>
        <DefinitionView definition={post} />
        <DialogFooter>
          <p className="text-muted-foreground text-xs">
            {new Date(post.createdAt).toLocaleDateString()}
          </p>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

export function PostViewSkeleton() {
  return (
    <Card>
      <CardContent>
        <DefinitionViewSkeleton />
      </CardContent>
    </Card>
  );
}
