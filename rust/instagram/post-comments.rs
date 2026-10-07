// How do I get comments on an Instagram post that is not mine?
// https://openhandle.dev/questions/how-do-i-get-comments-on-an-instagram-post-that-is-not-mine
// Run: cargo run --bin instagram-post-comments

use openhandle::Client;

#[tokio::main]
async fn main() -> Result<(), openhandle::Error> {
    let client = Client::new(std::env::var("OPENHANDLE_TEST_KEY").unwrap_or_default())?;

    // items() pages lazily, one request per page.
    let mut comments = client
        .instagram()
        .post("910100000001")
        .comments()
        .items(None);
    while let Some(comment) = comments.next().await {
        let comment = comment?;
        let author = comment.author.and_then(|author| author.handle);
        println!("{:?} {:?}", author, comment.text);
    }
    Ok(())
}
