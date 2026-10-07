// How do I search X (Twitter) posts by keyword?
// https://openhandle.dev/questions/how-do-i-search-x-twitter-posts-by-keyword
// Run: cargo run --bin twitter-search-posts

use openhandle::{Client, TwitterSearchPostsOptions};

#[tokio::main]
async fn main() -> Result<(), openhandle::Error> {
    let client = Client::new(std::env::var("OPENHANDLE_TEST_KEY").unwrap_or_default())?;

    let mut posts = client
        .twitter()
        .search()
        .posts()
        .items(TwitterSearchPostsOptions::new("synthetic").sort("latest"));
    while let Some(post) = posts.next().await {
        let post = post?;
        let author = post.author.and_then(|author| author.handle);
        println!("{:?} {:?}", author, post.text);
    }
    Ok(())
}
