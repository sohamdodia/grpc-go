package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sohamdodia/grpc-go/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Soham",
		Title:    "Dodia",
		Content:  "Content of the first blog",
	}
	createdBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})

	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	fmt.Printf("Blog has been created: %v\n", createdBlogRes)
	blogID := createdBlogRes.GetBlog().GetId()
	//Read Blog

	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "232wa"})

	if err2 != nil {
		fmt.Printf("Error happened while reading: %v\n", err)
	}

	readBlogRes, err3 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: blogID})
	if err3 != nil {
		fmt.Printf("Error happened while reading: %v \n", err)
	}

	fmt.Printf("Blog has been read: %v\n", readBlogRes)

	newBlog := &blogpb.Blog{
		Id:       createdBlogRes.GetBlog().GetId(),
		AuthorId: "Changed Author",
		Title:    "My first Blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})

	if updateErr != nil {
		fmt.Printf("Error happened while updating: %v \n", updateRes)
	}

	fmt.Printf("Blog was updated: %v\n", updateRes)

	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogID})
	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v \n", deleteErr)
	}

	fmt.Printf("Blog was deleted: %v\n", deleteRes)

}
