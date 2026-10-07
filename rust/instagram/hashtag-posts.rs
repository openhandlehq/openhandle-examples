// How do I get Instagram hashtag posts past the 30 hashtag limit?
// https://openhandle.dev/questions/how-do-i-get-instagram-hashtag-posts-past-the-30-hashtag-limit
// Run: cargo run --bin instagram-hashtag-posts

use openhandle::{Client, InstagramHashtagPostsOptions};

#[tokio::main]
async fn main() -> Result<(), openhandle::Error> {
    let client = Client::new(std::env::var("OPENHANDLE_TEST_KEY").unwrap_or_default())?;

    let mut posts = client
        .instagram()
        .hashtag("syntheticvolume")
        .posts()
        .items(InstagramHashtagPostsOptions::new().sort("recent"));
    while let Some(post) = posts.next().await {
        let post = post?;
        let author = post.author.and_then(|author| author.handle);
        println!("{} {:?} {:?}", post.id, author, post.metrics.likes);
    }
    Ok(())
}
