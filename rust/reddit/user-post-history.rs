// How do I get a Reddit user's post history?
// https://openhandle.dev/questions/how-do-i-get-a-reddit-users-post-history
// Run: cargo run --bin reddit-user-post-history

use openhandle::{Client, RedditProfilePostsOptions};

#[tokio::main]
async fn main() -> Result<(), openhandle::Error> {
    let client = Client::new(std::env::var("OPENHANDLE_TEST_KEY").unwrap_or_default())?;

    let mut posts = client
        .reddit()
        .profile("synthetic_reddit")
        .posts()
        .items(RedditProfilePostsOptions::new().sort("new"));
    while let Some(post) = posts.next().await {
        let post = post?;
        let community = post
            .community
            .and_then(|community| community.display_handle);
        println!("{:?} {} {:?}", community, post.created_at, post.title);
    }
    Ok(())
}
